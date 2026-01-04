package test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/DKeshavarz/eventic/pkg/utile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEvent(t *testing.T) {
	curTime := time.Now()

	validEvent := &entity.Event{
		OrganizerID: 7,
		Title:       "test",
		Cost:        100,
		DateTime:    utile.TimePtr(curTime.Add(12 * time.Hour)),
		Description: "some thing...",
	}

	validOrg := &entity.Organization{
		ID: 7,
		OwnerID: 5,
		Name: "my valid Organization",
	}

	testCases := []struct {
		name      string
		userID    int
		event     *entity.Event
		setupMock func(m *mockEventStorage, o *mockOrganizionStorage)
		wantEvent *entity.Event
		wantErr   error
	}{
		{
			name:  "Valid event",
			userID: 5,
			event: validEvent,
			setupMock: func(m *mockEventStorage, o *mockOrganizionStorage) {
				o.On("GetByID", validEvent.OrganizerID).Return(validOrg, nil)
				m.On("Create", mock.Anything).Return(validEvent, nil)
			},
			wantEvent: validEvent,
			wantErr:   nil,
		},
		{
			name: "Lost title",
			userID: 5,
			event: &entity.Event{
				Title:       "",
				Cost:        100,
				Description: "some thing...",
				DateTime:    utile.TimePtr(curTime.Add(12 * time.Hour)),
			},
			setupMock: func(m *mockEventStorage, o *mockOrganizionStorage) {},
			wantErr:   entity.ErrInvalidTitle,
		},
		{
			name: "Negetive cost",
			userID: 5,
			event: &entity.Event{
				Title:       "title",
				Cost:        -50,
				Description: "some thing...",
				DateTime:    utile.TimePtr(curTime.Add(12 * time.Hour)),
			},
			setupMock: func(m *mockEventStorage, o *mockOrganizionStorage) {},
			wantErr:   entity.ErrInvalidCost,
		},
		{
			name: "Nonexisting organizaion",
			userID: 7,
			event: validEvent,
			setupMock: func(m *mockEventStorage, o *mockOrganizionStorage) {
				o.On("GetByID", validEvent.OrganizerID).Return(validOrg, nil)
				m.On("Create", mock.Anything).Return(validEvent, nil)
			},
			wantErr: event.ErrInvalidEventCreator,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eventStorage := new(mockEventStorage)
			joinEventStorage := new(mockJoinEventStorage)
			OrgStorage := new(mockOrganizionStorage)
			tc.setupMock(eventStorage, OrgStorage)
			userStorage := new(mockUserStorage)
			service := event.NewService(eventStorage, joinEventStorage,OrgStorage, userStorage)
			event, err := service.Create(tc.userID, tc.event)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantEvent, event)
		})
	}
}
