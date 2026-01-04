package event

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type Service interface {
	Create(userID int, event *entity.Event) (*entity.Event, error)
	Join(joinEvent *entity.JoinEvent) (*entity.JoinEvent, error)
	GetAll() ([]*entity.Event, error)
	Get(id int) (*entity.Event, error)
}

var (
	ErrInvalidUser         = errors.New("invalid user")
	ErrInvalidEvent        = errors.New("invalid event")
	ErrInvalidEventCreator = errors.New("This user can't create eve't in this organizaion")
)

type service struct {
	eventStorage        repositories.Event
	joinEventStorage    repositories.JoinEvent
	organizationStorage repositories.Organization
	userStorage         repositories.User
}

func NewService(eventStorage repositories.Event, joinEventStorage repositories.JoinEvent, org repositories.Organization, userStorage repositories.User) Service {
	return &service{
		eventStorage:        eventStorage,
		joinEventStorage:    joinEventStorage,
		organizationStorage: org,
		userStorage: userStorage,
	}
}
