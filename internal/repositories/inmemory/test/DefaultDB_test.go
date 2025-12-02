package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/DKeshavarz/eventic/pkg/utiles"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetByEmailWithDeafaultDB(t *testing.T) {
	db := inmemory.DefaultDB()
	userStore := inmemory.NewUserStorage(db)
	
	user1 := &entity.User{
		Username: "ali",
		Password: "1234",
		Email:    utiles.StrPtr("ali@example.com"),
	}
	newUser1, err := userStore.Create(user1)

	assert.Nil(t, err)
	assert.Equal(t, newUser1, user1)

	getUser, err := userStore.GetUserByEmail("ali@example.com")
	assert.Nil(t, err)
	assert.Equal(t, newUser1, getUser)
}
