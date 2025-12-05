package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type joinEventStorage struct {
	mock.Mock
}

func TestJoinEvent(t *testing.T) {
	testCases := []struct {
		name          string
		joinEvent     *entity.JoinEvent
		setupMock     func(m *joinEventStorage)
		wantErr       error
		wantJoinEvent *entity.JoinEvent
	}{
		{
			name: "Valid join request",
			joinEvent: &entity.JoinEvent{
				UserID: 12,
				EventID: 50,
			},
			setupMock: func(m *joinEventStorage) {
				m.On("Create", mock.Anything).Return(&entity.JoinEvent{
					UserID: 12,
					EventID: 50,
				}, nil)
			},
			wantErr: nil,
			wantJoinEvent: &entity.JoinEvent{
				UserID: 12,
				EventID: 50,
			},
		},
		{
			name: "Invalid join request _ invalid user",
			joinEvent: &entity.JoinEvent{
				UserID: 10,
				EventID: 51,
			},
			setupMock: func(m *joinEventStorage) {
				m.On("Create", mock.Anything).Return(&entity.JoinEvent{}, event.ErrInvalidUser)
			},
			wantErr: event.ErrInvalidUser,
			wantJoinEvent:  &entity.JoinEvent{},
		},
		{
			name: "Invalid join request _ invalid event",
			joinEvent: &entity.JoinEvent{
				UserID: 10,
				EventID: 510,
			},
			setupMock: func(m *joinEventStorage) {
				m.On("Create", mock.Anything).Return(&entity.JoinEvent{}, event.ErrInvalidEvent)
			},
			wantErr: event.ErrInvalidEvent,
			wantJoinEvent:  &entity.JoinEvent{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eventStorage := &eventStorage{}
			joinEventStorage := new(joinEventStorage)
			tc.setupMock(joinEventStorage)

			service := event.NewService(eventStorage, joinEventStorage)
			

			joinEvent, err := service.Join(tc.joinEvent)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantJoinEvent, joinEvent)

		})
	}
}

// -------------- helpers ----------------------
func (e *joinEventStorage) Create(org *entity.JoinEvent) (*entity.JoinEvent, error) {
	args := e.Called(org)
	return args.Get(0).(*entity.JoinEvent), args.Error(1)
}
