package service

import "KitaSehat_Backend/internal/features/users"

type UserService struct {
	qry users.UQuery
}

func NewUserService(q users.UQuery) users.UService {
	return &UserService{
		qry: q,
	}
}

func (us *UserService) Login(username string, password string) (users.User, error) {
	return users.User{}, nil
}

func (us *UserService) Register(input users.User) error {
	return nil
}

func (us *UserService) GetAllUsers() ([]users.User, error) {
	return []users.User{}, nil
}

func (us *UserService) GetUser(id int) (users.User, error) {
	return users.User{}, nil
}

func (us *UserService) UpdateUser(input users.User) error {
	return nil
}

func (us *UserService) DeleteUser(id int) error {
	return nil
}

func (us *UserService) ChangePassword(id int, password string) error {
	return nil
}
