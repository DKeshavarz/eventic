package usecase

import "github.com/DKeshavarz/eventic/internal/entity"

type User interface {
	JoinEvent(userID int,eventID int) error
	CreateOrganization(org *entity.Organization) (*entity.Organization, error) 
}