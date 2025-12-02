package test

import (
	"testing"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestJoinEventGetByUserID(t *testing.T) {
	db := inmemory.NewDB()
	JoinEventStorage := inmemory.NewJoinEventStorage(db)

	_, err := JoinEventStorage.GetByUserID(1)
	assert.Nil(t, err)
}

func TestCreateJoinEvent(t *testing.T) {
	db := inmemory.NewDB()
	JoinEventStorage := inmemory.NewJoinEventStorage(db)

	joinEvent := &entity.JoinEvent{
		EventID: 50,
		UserID:  1,
	}
	newJoinevent, err := JoinEventStorage.Create(joinEvent)
	assert.Nil(t, err)
	assert.Equal(t, newJoinevent, joinEvent)
}

func TestUserJoinEventAfterInsersion(t *testing.T) {
	db := inmemory.NewDB()
	JoinEventStorage := inmemory.NewJoinEventStorage(db)

	joinEventUser1 := []*entity.JoinEvent{
		{
			EventID: 50,
			UserID:  1,
		},
		{
			EventID: 52,
			UserID:  1,
		},
		{
			EventID: 53,
			UserID:  1,
		},
	}
	

	for _, value := range joinEventUser1 {
		new, err :=JoinEventStorage.Create(value)
		assert.Nil(t, err)
		assert.Equal(t, new, value)
	}

	list, err := JoinEventStorage.GetByUserID(1)
	assert.Nil(t, err)
	assert.ElementsMatch(t, joinEventUser1, list)

}
