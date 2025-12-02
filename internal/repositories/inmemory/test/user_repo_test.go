package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/DKeshavarz/eventic/pkg/utiles"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db := inmemory.NewDB()
	userStore := inmemory.NewUserStorage(db)

	user1 := &entity.User{
		Username: "ali",
		Password: "1234",
		Email:    utiles.StrPtr("ali@example.com"),
	}
	newUser1, err := userStore.Create(user1)

	assert.Nil(t, err)
	assert.Equal(t, newUser1, user1)


	user2 := &entity.User{
		Username: "reza",
		Password: "1234",
		Phone:    utiles.StrPtr("09398116589"),
	}
	newUser2, err := userStore.Create(user2)

	assert.Nil(t, err)
	assert.Equal(t, newUser2, user2)
	assert.NotEqual(t, newUser1, newUser2)
}

func TestCreateAndGetByEmail(t *testing.T) {
	db := inmemory.NewDB()
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

func TestCreateAndGetByPhone(t *testing.T) {
	db := inmemory.NewDB()
	userStore := inmemory.NewUserStorage(db)

	user := &entity.User{
		Username: "ali",
		Password: "1234",
		Phone:    utiles.StrPtr("09398116589"),
	}
	newUser, err := userStore.Create(user)

	assert.Nil(t, err)
	assert.Equal(t, newUser, user)

	_, err = userStore.GetUserByEmail("ali@example.com")
	if assert.NotNil(t, err) {
		assert.Equal(t, err, repositories.ErrUserNotFound)
	}

	getUser, err := userStore.GetUserByPhone("09398116589")
	assert.Nil(t, err)
	assert.Equal(t, newUser, getUser)
}
