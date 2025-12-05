package organization

import "github.com/DKeshavarz/eventic/internal/entity"

func (s *service) Create(org *entity.Organization) (*entity.Organization, error) {
	if err := org.Validate(); err != nil {
		return nil, err
	}

	newOrg, err := s.orgStorage.Create(org)

	if err != nil {
		return nil, err
	}
	return newOrg, nil
}
