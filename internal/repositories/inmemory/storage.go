package inmemory

import (
	"sync"

	"github.com/DKeshavarz/eventic/internal/entity"
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
