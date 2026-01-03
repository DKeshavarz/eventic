package test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEvent(t *testing.T) {
	curTime := time.Now()

	validEvent := &entity.Event{
		Title:       "test",
		Cost:        100,
		DateTime:    curTime.Add(12 * time.Hour),
		Description: "some thing...",
	}

	testCases := []struct {
		name      string
		userID    int
		event     *entity.Event
		setupMock func(m *mockEventStorage)
		wantEvent *entity.Event
		wantErr   error
	}{
		{
			name:  "valid event",
			userID: 5,
			event: validEvent,
			setupMock: func(m *mockEventStorage) {
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
				DateTime:    curTime.Add(12 * time.Hour),
			},
			setupMock: func(m *mockEventStorage) {},
			wantErr:   entity.ErrInvalidTitle,
		},
		{
			name: "Negetive cost",
			userID: 5,
			event: &entity.Event{
				Title:       "title",
				Cost:        -50,
				Description: "some thing...",
				DateTime:    curTime.Add(12 * time.Hour),
			},
			setupMock: func(m *mockEventStorage) {},
			wantErr:   entity.ErrInvalidCost,
		},
		// {
		// 	name: "Invalid user request",
		// 	userID: 7,
		// 	event: validEvent,
		// 	setupMock: func(m *mockEventStorage) {
		// 		m.On("Create", mock.Anything).Return(validEvent, nil)
		// 	},
		// 	wantErr: event.ErrInvalidEventCreator,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eventStorage := new(mockEventStorage)
			joinEventStorage := new(mockJoinEventStorage)
			OrgStorage := new(mockOrganizionStorage)
			tc.setupMock(eventStorage)
			service := event.NewService(eventStorage, joinEventStorage,OrgStorage)
			event, err := service.Create(tc.userID, tc.event)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantEvent, event)
		})
	}
}
