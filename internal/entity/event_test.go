package entity

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/pkg/utile"
)

func TestValidateEvent(t *testing.T) {
	testCases := []struct {
		name        string
		event       *Event
		expectedErr error
	}{
		{
			name: "Valid Event",
			event: &Event{
				ID:          1,
				OrganizerID: 1,
				Title:       "Test Event",
				Cost:        100,
				DateTime:    time.Now(),
				Description: "This is a test event",
				Location:    nil,
				PosterPic:   nil,
				Link:        nil,
			},
			expectedErr: nil,
		},
		{
			name: "invalid Event - negetive cost",
			event: &Event{
				ID:          2,
				OrganizerID: 1,
				Title:       "Free Event",
				Cost:        -10,
				DateTime:    time.Now(),
				Description: "This is a free event",
				Location:    nil,
				PosterPic:   nil,
				Link:        nil,
			},
			expectedErr: ErrInvalidCost,
		},
		{
			name: "Invalid Event - empty title",
			event: &Event{
				ID:          3,
				OrganizerID: 1,
				Title:       "",
				Cost:        100,
				DateTime:    time.Now(),
				Description: "This is a test event",
				Location:    nil,
				PosterPic:   nil,
				Link:        nil,
			},
			expectedErr: ErrInvalidTitle,
		},
		{
			name: "Invalid Event - empty description",
			event: &Event{
				ID:          5,
				OrganizerID: 1,
				Title:       "Test Event",
				Cost:        100,
				DateTime:    time.Now(),
				Description: "",
				Location:    nil,
				PosterPic:   nil,
				Link:        nil,
			},
			expectedErr: ErrInvalidDescription,
		},
		{
			name: "Valid Event - with all optional fields",
			event: &Event{
				ID:          6,
				OrganizerID: 1,
				Title:       "Full Event",
				Cost:        200,
				DateTime:    time.Now(),
				Description: "This is a complete event",
				Location:    utile.StrPtr("123 Main St"),
				PosterPic:   utile.StrPtr("/images/poster.jpg"),
				Link:        utile.StrPtr("https://example.com/event"),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.event.Validate()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}


