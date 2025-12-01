package repositories

import "github.com/DKeshavarz/eventic/internal/entity"

type User interface {
	GetUserByPhone(phone string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}	

type Organization interface {
	Create(org *entity.Organization) (*entity.Organization, error)
}

type Event interface {
	Create(event *entity.Event)(*entity.Event, error)
}

type JoinEvent interface {
	Create(event *entity.JoinEvent)(*entity.JoinEvent, error)
}