package entity

import (
	"errors"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	OrganizerID int       `json:"organizer_id"`
	Title       string    `json:"title"`
	Cost        int       `json:"cost"`
	DateTime    time.Time `json:"datetime"`
	Description string    `json:"description"`
	Location    *string   `json:"location"`
	PosterPic   *string   `json:"poster_pic"`
	Link        *string   `json:"link"`
}

var (
	ErrInvalidTitle       = errors.New("invalid title")
	ErrInvalidCost        = errors.New("invalid cost")
	ErrInvalidDateTime    = errors.New("invalid datetime")
	//ErrInvalidDescription = errors.New("invalid description")
)

func (e *Event) Validate() error {
	if len(e.Title) == 0 {
		return ErrInvalidTitle
	}

	if e.Cost < 0 {
		return ErrInvalidCost
	}

	if len(e.Description) == 0 {
		return ErrInvalidDescription
	}

	return nil
}

type JoinEvent struct {
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}
