package repositories

import (
	"errors"

	"github.com/DKeshavarz/eventic/internal/entity"
)

var (
	ErrUserNotFound  = errors.New("کاربر پیدا نشد")
	ErrEventNotFound = errors.New("رویداد پیدا نشد")
	ErrOrgNotFound   = errors.New("سازمان مورد نظر یافت نشد")
)

type User interface {
	Create(user *entity.User) (*entity.User, error)
	GetUserByPhone(phone string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetByID(id int) (*entity.User, error)
}

type Organization interface {
	GetByID(id int) (*entity.Organization, error)
	Create(org *entity.Organization) (*entity.Organization, error)
}

type Event interface {
	GetByID(id int) (*entity.Event, error)
	Create(event *entity.Event) (*entity.Event, error)
	GetAll() ([]*entity.Event, error)
}

type JoinEvent interface {
	GetByUserID(id int) ([]*entity.JoinEvent, error)
	Create(event *entity.JoinEvent) (*entity.JoinEvent, error)
}
