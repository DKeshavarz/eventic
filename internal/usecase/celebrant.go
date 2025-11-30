package usecase

import "github.com/DKeshavarz/eventic/internal/entity"


type Celebrant interface {
	CreateEvent() (entity.Event, error)
} 