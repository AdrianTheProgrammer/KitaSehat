package repository

import (
	"KitaSehat_Backend/internal/features/users"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(DBConnect *gorm.DB) users.UQuery {
	return &UserQuery{
		db: DBConnect,
	}
}

func (uq *UserQuery) Login(username string) (users.User, error) {
	return users.User{}, nil
}

func (uq *UserQuery) Register(input users.User) error {
	return nil
}

func (uq *UserQuery) GetAllUsers() ([]users.User, error) {
	return []users.User{}, nil
}

func (uq *UserQuery) GetUser(id int) (users.User, error) {
	return users.User{}, nil
}

func (uq *UserQuery) UpdateUser(input users.User) error {
	return nil
}

func (uq *UserQuery) DeleteUser(id int) error {
	return nil
}

func (uq *UserQuery) ChangePassword(id int, password string) error {
	return nil
}
