package test

import (
	"testing"
	"time"

	"github.com/DKeshavarz/eventic/internal/repositories"
	"github.com/DKeshavarz/eventic/internal/repositories/cache"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	storage := cache.New()
	_, err := storage.Get("Ali")
	assert.Equal(t, err, repositories.ErrNotFoundOTP)
}

func TestSetThenGet(t *testing.T) {
	storage := cache.New()
	err := storage.Set("key", "value", 1 * time.Second) 
	assert.Nil(t, err)

	value, err := storage.Get("key")
	assert.Nil(t, err)
	assert.Equal(t, value, "value")
}

func TestSetDeleteGet(t *testing.T) {
	storage := cache.New()
	err := storage.Set("key", "value", 1 * time.Second) 
	assert.Nil(t, err)

	err = storage.Delete("key")	
	assert.Nil(t, err)

	_, err = storage.Get("key")
	assert.Equal(t, err, repositories.ErrNotFoundOTP)
}

func TestDelete(t *testing.T) {
	storage := cache.New()
	err := storage.Delete("key")	
	assert.Equal(t, err, repositories.ErrDeleteOTP)
}

func TestGetWithTll(t *testing.T) {
	storage := cache.New()
	err := storage.Set("key", "value", 1 * time.Second) 
	assert.Nil(t, err)

	time.Sleep(2 * time.Second)

	_, err = storage.Get("key")
	assert.Equal(t, err, repositories.ErrNotFoundOTP)
}
