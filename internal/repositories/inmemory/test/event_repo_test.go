package test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/repositories/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestEventGetByID(t *testing.T) {
	db := inmemory.NewDB()
	eventStorage := inmemory.NewEventStorage(db)

	_, err := eventStorage.GetByID(1)
	if assert.NotNil(t, err) {
		assert.Equal(t, err, repositories.ErrEventNotFound)
	}
}
func TestCreateEvent(t *testing.T) {
	db := inmemory.NewDB()
	eventStorage := inmemory.NewEventStorage(db)

	event := &entity.Event{
		Title: "title",
		Description: "des",
		Cost: 100,
		DateTime: time.Now().Add(time.Hour * 72),
	}

	newEvent, err := eventStorage.Create(event)
	assert.Nil(t, err)
	assert.Equal(t, newEvent, event)
}

func TestCreateAndGet(t *testing.T){
	db := inmemory.NewDB()
	eventStorage := inmemory.NewEventStorage(db)

	event := &entity.Event{
		Title: "title",
		Description: "des",
		Cost: 100,
		DateTime: time.Now().Add(time.Hour * 72),
	}

	newEvent, err := eventStorage.Create(event)
	assert.Nil(t, err)
	assert.Equal(t, newEvent, event)

	storeEvent , err := eventStorage.GetByID(newEvent.ID)
	assert.Nil(t, err)
	assert.Equal(t, storeEvent, event)
}

func TestCreateAndGetAll(t *testing.T){
	db := inmemory.NewDB()
	eventStorage := inmemory.NewEventStorage(db)

	event1 := &entity.Event{
		Title: "title",
		Description: "descripton",
		Cost: 100,
		DateTime: time.Now().Add(time.Hour * 72),
	}
	newEvent, err := eventStorage.Create(event1)
	assert.Nil(t, err)
	assert.Equal(t, newEvent, event1)

	event2 := &entity.Event{
		Title: "title 2",
		Description: "descripton 2",
		Cost: 1000,
		DateTime: time.Now().Add(time.Hour * 2),
	}
	newEvent, err = eventStorage.Create(event2)
	assert.Nil(t, err)
	assert.Equal(t, newEvent, event2)

	allEvents, err := eventStorage.GetAll()
	assert.Nil(t, err)
	assert.ElementsMatch(t, allEvents, []*entity.Event{event1, event2})
}
