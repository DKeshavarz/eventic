package auth

import (
	"fmt"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity/validation"
	"github.com/DKeshavarz/eventic/internal/getways"
)

func (s *service) SendOTP(mail string, expire time.Duration) error {
	if expire <= 0 {
		return ErrInvalidExpire
	}
	if err := validation.Email(mail); err != nil {
		return err
	}
	
	code, err := generateCode(6)
	if err != nil {
		return err
	}
	
	err = s.cache.Set(mail, code, expire)
	if err != nil {
		return err
	}
	
	err = s.sender.Send(mail, &getways.Message{
		Title: "This is you otp Code",
		Text: fmt.Sprintf("Use this %s Code to verify your acount", code),
	})
	if err != nil {
		return err
	}
	
	return nil
}
func (s *service) VerifyOTP(mail string, code string) error {
	realCode, err := s.cache.Get(mail)
	if err != nil {
		return err
	}
	if code != realCode {
		return ErrWrongCode
	}
	return nil
}

func (s *service)getCode(mail string) string {
	code, err := s.cache.Get(mail)
	if err != nil {
		return ""
	}
	return code
}
