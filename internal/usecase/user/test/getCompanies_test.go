package usecase

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/usecase/user"
	"github.com/stretchr/testify/assert"
)

func TestGetCompanies(t *testing.T) {
	tests := []struct {
		tag               string
		id                int
		setupMock         func(m *mockUserStorage, o *mockOrgStorage)
		wantOrganizations []*entity.Organization
		wantErr           error
	}{
		{
			tag:               "User with no company",
			id:                16,
			setupMock:         func(m *mockUserStorage, o *mockOrgStorage) {
				m.On("GetByID", 16).Return(&entity.User{ID: 16}, nil)
				o.On("GetByOwnerID", 16).Return([]*entity.Organization{}, nil)
			},
			wantOrganizations: []*entity.Organization{},
		},
		{
			tag: "Nonexisting user",
			id:  17,
			setupMock: func(m *mockUserStorage, o *mockOrgStorage) {
				m.On("GetByID", 17).Return(nil, repositories.ErrUserNotFound)
			},
			wantErr: repositories.ErrUserNotFound,
		},
		{
			tag: "User with comapnies",
			id:  18,
			setupMock: func(m *mockUserStorage, o *mockOrgStorage) {
				m.On("GetByID", 18).Return(&entity.User{ID: 18}, nil)
				o.On("GetByOwnerID", 18).Return([]*entity.Organization{
					{ID: 1, OwnerID: 18, Name: "name 1"},
					{ID: 4, OwnerID: 18, Name: "name 2"},
				}, nil)
			},
			wantOrganizations: []*entity.Organization{
				{ID: 1, OwnerID: 18, Name: "name 1"},
				{ID: 4, OwnerID: 18, Name: "name 2"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {
			userStorage := new(mockUserStorage)
			orgStorage := new(mockOrgStorage)
			tc.setupMock(userStorage, orgStorage)
			service := user.NewSevice(userStorage, orgStorage)
			orgs, err := service.GetCompanies(tc.id)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.ElementsMatch(t, orgs, tc.wantOrganizations)
			assert.Nil(t, err)
		})
	}
}
