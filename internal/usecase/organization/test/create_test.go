package organization

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/usecase/organization"
	"github.com/DKeshavarz/eventic/pkg/utiles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type organizationStorage struct {
	mock.Mock
}

func TestCreateOrg(t *testing.T) {
	testCases := []struct {
		name      string
		org       *entity.Organization
		setupMock func(m *organizationStorage)
		wantErr   error
		wantOrg   *entity.Organization
	}{
		{
			name: "valid organization",
			org: &entity.Organization{
				OwnerID:     5,
				Name:        "test",
				Description: "somthing",
			},
			setupMock: func(m *organizationStorage) {
				m.On("Create", mock.Anything).Return(&entity.Organization{
					ID:          7,
					OwnerID:     5,
					Name:        "test",
					Description: "somthing",
				}, nil)
			},
			wantErr: nil,
			wantOrg: &entity.Organization{
				ID:          7,
				OwnerID:     5,
				Name:        "test",
				Description: "somthing",
			},
		},
		{
			name: "Invalid organization - missing descitption",
			org: &entity.Organization{
				OwnerID:     5,
				Name:        "test",
				Description: "",
			},
			setupMock: func(m *organizationStorage) {
				m.On("Create", mock.Anything).Return(&entity.Organization{}, nil)
			},
			wantErr: entity.ErrInvalidDescription,
			wantOrg: nil,
		},
		{
			name: "Invalid organization - Bad Email",
			org: &entity.Organization{
				OwnerID:     5,
				Name:        "test",
				Description: "somthing",
				Email:       utiles.StrPtr("somebadmal+e@sss.com"),
			},
			setupMock: func(m *organizationStorage) {
				m.On("Create", mock.Anything).Return(&entity.Organization{}, nil)
			},
			wantErr: entity.ErrInvalidEmail,
			wantOrg: nil,
		},
		{
			name: "Invalid organization - Not existing owner",
			org: &entity.Organization{
				OwnerID:     10,
				Name:        "test name",
				Description: "somthing",
				Email:       utiles.StrPtr("danny@sss.com"),
			},
			setupMock: func(m *organizationStorage) {
				m.On("Create", mock.Anything).Return(&entity.Organization{}, organization.ErrInvalidOwner)
			},
			wantErr: organization.ErrInvalidOwner,
			wantOrg: nil,
		},
		{
			name: "Invalid organization - Douplicated name",
			org: &entity.Organization{
				OwnerID:     10,
				Name:        "test name",
				Description: "somthing",
				Email:       utiles.StrPtr("danny@sss.com"),
			},
			setupMock: func(m *organizationStorage) {
				m.On("Create", mock.Anything).Return(&entity.Organization{}, organization.ErrDuplicatedName)
			},
			wantErr: organization.ErrDuplicatedName,
			wantOrg: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			orgStorage := new(organizationStorage)
			tc.setupMock(orgStorage)

			orgService := organization.NewService(orgStorage)

			org, err := orgService.Create(tc.org)

			if tc.wantErr != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantOrg, org)
		})
	}

}

// --------------- helpers -----------------------

func (m *organizationStorage) Create(org *entity.Organization) (*entity.Organization, error) {
	args := m.Called(org)
	return args.Get(0).(*entity.Organization), args.Error(1)
}
