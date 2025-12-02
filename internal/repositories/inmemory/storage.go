package inmemory

import (
	"sync"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/pkg/utiles"
)

type DB struct {
	users         map[int]*entity.User
	organizations map[int]*entity.Organization
	events        map[int]*entity.Event
	joinEvents    map[string]*entity.JoinEvent

	mu           sync.RWMutex
	userCounter  int
	orgCounter   int
	eventCounter int
}

func NewDB() *DB {
	return &DB{
		users:         make(map[int]*entity.User),
		organizations: make(map[int]*entity.Organization),
		events:        make(map[int]*entity.Event),
		joinEvents:    make(map[string]*entity.JoinEvent),
		userCounter:   1,
		orgCounter:    1,
		eventCounter:  1,
	}
}

func DefaultDB() *DB {
	users := map[int]*entity.User{
		1: {
			ID:       1,
			Username: "Danny",
			Password: "1234",
			Email:    utiles.StrPtr("dankeshavarz1075@gmail.com"),
		},
	}
	organizations := make(map[int]*entity.Organization)
	events := make(map[int]*entity.Event)
	joinEvents := make(map[string]*entity.JoinEvent)

	return &DB{
		users:         users,
		organizations: organizations,
		events:        events,
		joinEvents:    joinEvents,
		userCounter:   len(users) + 1,
		orgCounter:    len(organizations) + 1,
		eventCounter:  len(events) + 1,
	}
}
