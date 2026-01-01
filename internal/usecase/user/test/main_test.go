package usecase

import "testing"

var (
	userStorage *mockUserStorage
)

func TestMain(m *testing.M) {
	userStorage = new(mockUserStorage)
}
