package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/getways"
	"github.com/DKeshavarz/eventic/internal/repositories/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendOTP(t *testing.T) {
	cache := cache.New()
	sender := &mockSender{}
	sender.On("Send", "GoodMail@gmail.com").Return(
		nil,
	)
	service := New(cache, sender)

	err := service.SendOTP("BadMail@gmail.c", time.Minute*5)
	assert.NotNil(t, err)

	err = service.SendOTP("BadMail@gmail.c", -time.Minute*5)
	assert.Equal(t, ErrInvalidExpire, err)

	err = service.SendOTP("GoodMail@gmail.com", time.Minute*5)
	assert.Nil(t, err)
}

func TestVerify(t *testing.T) {
	cache := cache.New()
	var sender getways.Sender
	service := New(cache, sender)

	err := service.VerifyOTP("NonExistingMail@gmail.com", "some Code")

	assert.NotNil(t, ErrWrongCode, err)
}

func TestGetAndVerify(t *testing.T) {
	cache := cache.New()
	sender := &mockSender{}
	sender.On("Send", "SomeMail@gmail.com").Return(
		nil,
	)
	service := New(cache, sender)

	err := service.SendOTP("SomeMail@gmail.com", time.Minute*5)
	assert.Nil(t, err)
	badCode := ""
	err = service.VerifyOTP("SomeMail@gmail.com", badCode)
	assert.Equal(t, ErrWrongCode, err)

	realCode := service.getCode("SomeMail@gmail.com")
	err = service.VerifyOTP("SomeMail@gmail.com", realCode)
	assert.Nil(t, err)
}

func TestSendOTPEmail(t *testing.T) {
	message := &getways.Message{
		Title: "test",
		Text:  "body",
	}
	wantErr := errors.New("some error")

	cache := cache.New()
	sender := &mockSender{}
	sender.On("Send", "SomeMail@gmail.com", message).Return(
		wantErr,
	)
	sender.On("Send", "GoodMail@gmail.com", message).Return(
		nil,
	)

	service := New(cache, sender)

	err := service.SendOTP("SomeMail@gmail.com", time.Second*1)
	assert.Equal(t, err, wantErr)

	err = service.SendOTP("GoodMail@gmail.com", time.Second*1)
	assert.Nil(t, err)

	code := service.getCode("GoodMail@gmail.com")
	err = service.VerifyOTP("GoodMail@gmail.com", code)
	assert.Nil(t, err)

	time.Sleep(time.Second * 2)

	err = service.VerifyOTP("GoodMail@gmail.com", code)
	assert.NotNil(t, err)
}

func TestCache(t *testing.T) {
	cache := cache.New()

	err := cache.Set("key", "value", 1*time.Second)
	assert.Nil(t, err)

	value, err := cache.Get("key")
	assert.Equal(t, value, "value")
	assert.Nil(t, err)

	time.Sleep(2 * time.Second)

	_, err = cache.Get("key")
	assert.NotNil(t, err)
}

// ------------ mocks -----------------------
type mockSender struct {
	mock.Mock
}

func (s *mockSender) Send(to string, msg *getways.Message) error {
	args := s.Called(to, mock.Anything)
	return args.Error(0)
}
