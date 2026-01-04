package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestOrgGetByID(t *testing.T) {
	db := inmemory.NewDB()
	orgStorage := inmemory.NewOrgStorage(db)

	_, err := orgStorage.GetByID(1)
	if assert.NotNil(t, err) {
		assert.Equal(t, err, repositories.ErrOrgNotFound)
	}
}

func TestCreateOrg(t *testing.T) {
	db := inmemory.NewDB()
	orgStorage := inmemory.NewOrgStorage(db)

	org := &entity.Organization{
		OwnerID: 1,
		Name: "org name",
		Description: "Some time's I feel I don't have courage to do the right thing",
	}

	newEvent, err := orgStorage.Create(org)
	assert.Nil(t, err)
	assert.Equal(t, newEvent, org)
}

func TestCreateAndGetOrg(t *testing.T){
	db := inmemory.NewDB()
	orgStorage := inmemory.NewOrgStorage(db)

	org := &entity.Organization{
		OwnerID: 1,
		Name: "org name",
		Description: "Some time's I feel I don't have courage to do the right thing",
	}

	newOrg, err := orgStorage.Create(org)
	assert.Nil(t, err)
	assert.Equal(t, newOrg, org)

	storeorg , err := orgStorage.GetByID(newOrg.ID)
	assert.Nil(t, err)
	assert.Equal(t, storeorg, org)
}

func TestOrgGetByOwnerID(t *testing.T) {
	db := inmemory.NewDB()
	orgStorage := inmemory.NewOrgStorage(db)

	data := make([]*entity.Organization,0)

	organizations, err := orgStorage.GetByOwnerID(12)
	assert.Nil(t, err)
	assert.ElementsMatch(t, data, organizations)
}

func TestOrgGetByOwnerIDWithData(t *testing.T) {
	db := inmemory.NewDB()
	orgStorage := inmemory.NewOrgStorage(db)

	data := []*entity.Organization{
		{OwnerID: 5},
		{OwnerID: 5},
		{OwnerID: 5},
		{OwnerID: 2},
	}
	for i := range data {
		new, err := orgStorage.Create(data[i])
		assert.Nil(t, err)
		assert.Equal(t, new, data[i])
	}
	
	organizations, err := orgStorage.GetByOwnerID(5)
	assert.Nil(t, err)
	assert.ElementsMatch(t, organizations, data[:3])

	organizations, err = orgStorage.GetByOwnerID(2)
	assert.Nil(t, err)
	assert.ElementsMatch(t, organizations, data[3:])
}