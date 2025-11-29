package entity

import "time"

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

type JoinEvent struct {
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}
