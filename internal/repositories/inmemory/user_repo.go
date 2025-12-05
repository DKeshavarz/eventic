package inmemory

import (
	"github.com/DKeshavarz/eventic/internal/entity"
	"github.com/DKeshavarz/eventic/internal/repositories"
)

type UserStorage struct {
	db *DB
}

func NewUserStorage(db *DB) repositories.User {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) GetUserByPhone(phone string) (*entity.User, error) {
	u.db.mu.RLock()
	defer u.db.mu.RUnlock()

	for _, val := range u.db.users {
		if val.Phone != nil && phone == *val.Phone {
			return val, nil
		}
	}

	return nil, repositories.ErrUserNotFound

}
func (u *UserStorage) GetUserByEmail(email string) (*entity.User, error) {
	u.db.mu.RLock()
	defer u.db.mu.RUnlock()

	for _, val := range u.db.users {
		if val.Email != nil && *(val.Email) == email {
			return val, nil
		}
	}

	return nil, repositories.ErrUserNotFound

}
func (u *UserStorage) Create(user *entity.User) (*entity.User, error) {
	u.db.mu.Lock()
	defer u.db.mu.Unlock()

	user.ID = u.db.userCounter
	u.db.users[u.db.userCounter] = user
	u.db.userCounter++
	return user, nil
}

func (u *UserStorage) GetByID(id int) (*entity.User, error) {
	if user, exist := u.db.users[id]; exist {
		return  user, nil
	}
	return nil, repositories.ErrUserNotFound
}


