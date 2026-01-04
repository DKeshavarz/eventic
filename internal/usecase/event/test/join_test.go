package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestJoinEvent(t *testing.T) {
	testCases := []struct {
		name          string
		joinEvent     *entity.JoinEvent
		setupMock     func(m *mockJoinEventStorage, u *mockUserStorage, e *mockEventStorage)
		wantErr       error
		wantJoinEvent *entity.JoinEvent
	}{
		{
			name: "Valid join request",
			joinEvent: &entity.JoinEvent{
				UserID: 12,
				EventID: 50,
			},
			setupMock: func(m *mockJoinEventStorage, u *mockUserStorage, e *mockEventStorage) {
				u.On("GetByID", 12).Return(&entity.User{ID: 12}, nil)
				e.On("GetByID", 50).Return(&entity.Event{ID: 50}, nil)
				m.On("Create", mock.Anything).Return(&entity.JoinEvent{UserID: 12,EventID: 50,}, nil)
			},
			wantErr: nil,
			wantJoinEvent: &entity.JoinEvent{
				UserID: 12,
				EventID: 50,
			},
		},
		{
			name: "Nonexiting  user",
			joinEvent: &entity.JoinEvent{
				UserID: 10,
				EventID: 51,
			},
			setupMock: func(m *mockJoinEventStorage, u *mockUserStorage, e *mockEventStorage) {
				u.On("GetByID", 10).Return(&entity.User{}, repositories.ErrUserNotFound)
			},
			wantErr: repositories.ErrUserNotFound,
		},
		{
			name: "Nonexiting event",
			joinEvent: &entity.JoinEvent{
				UserID: 10,
				EventID: 51,
			},
			setupMock: func(m *mockJoinEventStorage, u *mockUserStorage, e *mockEventStorage) {
				u.On("GetByID", 10).Return(&entity.User{ID: 10}, nil)
				e.On("GetByID", 51).Return(&entity.Event{}, repositories.ErrEventNotFound )
			},
			wantErr: repositories.ErrEventNotFound,
			wantJoinEvent:  &entity.JoinEvent{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eventStorage := new(mockEventStorage)
			joinEventStorage := new(mockJoinEventStorage)
			OrgStorage := new(mockOrganizionStorage)
			userStorage := new(mockUserStorage)
			tc.setupMock(joinEventStorage, userStorage, eventStorage)

			service := event.NewService(eventStorage, joinEventStorage, OrgStorage, userStorage)
			

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
func (e *mockJoinEventStorage) Create(org *entity.JoinEvent) (*entity.JoinEvent, error) {
	args := e.Called(org)
	return args.Get(0).(*entity.JoinEvent), args.Error(1)
}
