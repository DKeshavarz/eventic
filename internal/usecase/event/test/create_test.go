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
	testCases := []struct {
		name      string
		event     *entity.Event
		setupMock func(m *mockEventStorage)
		wantEvent *entity.Event
		wantErr   error
	}{
		{
			name: "valid event",
			event: &entity.Event{
				Title:       "test",
				Cost:        100,
				DateTime:    curTime.Add(12 * time.Hour),
				Description: "some thing...",
			},
			setupMock: func(m *mockEventStorage) {
				m.On("Create", mock.Anything).Return(&entity.Event{
					ID:          1,
					Title:       "test",
					Cost:        100,
					DateTime:    curTime.Add(12 * time.Hour),
					Description: "some thing...",
				}, nil)
			},
			wantEvent: &entity.Event{
				ID:          1,
				Title:       "test",
				Cost:        100,
				DateTime:    curTime.Add(12 * time.Hour),
				Description: "some thing...",
			},
			wantErr: nil,
		},
		{
			name: "invalid event - lost title",
			event: &entity.Event{
				Title:       "",
				Cost:        100,
				Description: "some thing...",
				DateTime:    curTime.Add(12 * time.Hour),
			},
			setupMock: func(m *mockEventStorage) {
				m.On("Create", mock.Anything).Return(&entity.Event{}, nil)
			},
			wantEvent: &entity.Event{},
			wantErr:   entity.ErrInvalidTitle,
		},
		{
			name: "invalid event - negetive cost",
			event: &entity.Event{
				Title:       "title",
				Cost:        -50,
				Description: "some thing...",
				DateTime:    curTime.Add(12 * time.Hour),
			},
			setupMock: func(m *mockEventStorage) {
				m.On("Create", mock.Anything).Return(&entity.Event{}, nil)
			},
			wantEvent: &entity.Event{},
			wantErr:   entity.ErrInvalidCost,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eventStorage := new(mockEventStorage)
			joinEventStorage := new(mockJoinEventStorage)
			tc.setupMock(eventStorage)
			service := event.NewService(eventStorage, joinEventStorage)
			event, err := service.Create(tc.event)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantEvent, event)
		})
	}
}
