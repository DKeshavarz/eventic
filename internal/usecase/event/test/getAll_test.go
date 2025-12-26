package test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	curTime := time.Now()
	event1 := &entity.Event{
		Title:       "test",
		Cost:        100,
		DateTime:    curTime.Add(12 * time.Hour),
		Description: "some thing...",
	}
	event2 := &entity.Event{
		Title:       "test 2",
		Cost:        1000,
		DateTime:    curTime.Add(120 * time.Hour),
		Description: "some thing 2...",
	}
	events := []*entity.Event{event1, event2}

	testCases := []struct {
		tag       string
		setupMock func(m *eventStorage)
		wantEvent []*entity.Event
		wantErr   error
	}{
		{
			tag: "Valid GetAll with nothing",
			setupMock: func(m *eventStorage) {
				m.On("GetAll").Return([]*entity.Event{}, nil)
			},
			wantEvent: []*entity.Event{},
			wantErr:   nil,
		},
		{
			tag: "Valid GetAll with somthig",
			setupMock: func(m *eventStorage) {
				m.On("GetAll").Return(events, nil)
			},
			wantEvent: events,
			wantErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.tag, func(t *testing.T) {
			eventStorage := new(eventStorage)
			tc.setupMock(eventStorage)
			joinEventStorage := new(joinEventStorage)

			service := event.NewService(eventStorage, joinEventStorage)
			events, err := service.GetAll()

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			} else {
				assert.ElementsMatch(t, tc.wantEvent, events)
				assert.Equal(t, tc.wantErr, err)
			}

		})
	}

}
