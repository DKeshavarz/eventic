package test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/usecase/event"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	curTime := time.Now()
	event1 := &entity.Event{
		ID:          5,
		Title:       "test",
		Cost:        100,
		DateTime:    curTime.Add(12 * time.Hour),
		Description: "some thing...",
	}
	event2 := &entity.Event{
		ID:          7,
		Title:       "test 2",
		Cost:        1000,
		DateTime:    curTime.Add(120 * time.Hour),
		Description: "some thing 2...",
	}

	testCases := []struct {
		tag       string
		id        int
		setupMock func(m *eventStorage)
		wantEvent *entity.Event
		wantErr   error
	}{
		{
			tag: "Want existing event with id 5",
			id:  5,
			setupMock: func(m *eventStorage) {
				m.On("GetByID", 5).Return(event1, nil)
			},
			wantEvent: event1,
			wantErr:   nil,
		},
		{
			tag: "Want existing event with id 7",
			id:  7,
			setupMock: func(m *eventStorage) {
				m.On("GetByID", 7).Return(event2, nil)
			},
			wantEvent: event2,
			wantErr:   nil,
		},
		{
			tag: "Want non-existing event",
			id:  70,
			setupMock: func(m *eventStorage) {
				m.On("GetByID", 70).Return(new(entity.Event), repositories.ErrEventNotFound)
			},
			wantEvent: nil,
			wantErr:  repositories.ErrEventNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.tag, func(t *testing.T) {
			eventStorage := new(eventStorage)
			tc.setupMock(eventStorage)
			joinEventStorage := new(joinEventStorage)

			service := event.NewService(eventStorage, joinEventStorage)
			events, err := service.Get(tc.id)
			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			} else {
				assert.Equal(t, tc.wantEvent, events)
				assert.Equal(t, tc.wantErr, err)
			}
		})
	}

}
