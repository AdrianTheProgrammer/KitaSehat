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

func (uq *UserQuery) Login(email string) (users.User, error) {
	// Find User by Email
	var result User

	err := uq.db.Where("email = ?", email).First(&result).Error
	if err != nil {
		return users.User{}, err
	}

	return ToUserEntity(result), nil
}

func (uq *UserQuery) Register(input users.User) error {
	// Insert Into Database
	cnv := ToUserData(input)

	err := uq.db.Create(&cnv).Error
	if err != nil {
		return err
	}

	return nil
}

func (uq *UserQuery) IsEmailRegistered(email string) bool {
	// Check if Email already registered
	err := uq.db.Where("email = ?", email).First(&User{}).Error
	return err == nil
}

func (uq *UserQuery) GetAllUsers(currentPage int, limit int, offset int) ([]users.User, int, error) {
	// Get All User Data From Database
	var result []User
	var totalItems int64

	err := uq.db.Limit(limit).Offset(offset).Find(&result).Error
	if err != nil {
		return []users.User{}, 0, err
	}

	// Get Total Items Count
	err = uq.db.Model(&User{}).Count(&totalItems).Error
	if err != nil {
		return []users.User{}, 0, err
	}

	return ToAllUserEntity(result), int(totalItems), nil
}

func (uq *UserQuery) GetUser(id int) (users.User, error) {
	// Get User Data By ID
	var result User

	err := uq.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return users.User{}, err
	}

	return ToUserEntity(result), nil
}

func (uq *UserQuery) UpdateUser(id int, input users.User) error {
	// Update User Data By ID
	err := uq.db.Model(&User{}).Where("id = ?", id).Updates(ToUserData(input)).Error
	if err != nil {
		return err
	}

	return nil
}

func (uq *UserQuery) DeleteUser(id int) error {
	// Delete User Data By ID
	err := uq.db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (uq *UserQuery) ChangePassword(id int, input users.User) error {
	// Change User Password By ID
	err := uq.db.Model(&User{}).Where("id = ?", id).Update("password", input.Password).Error
	if err != nil {
		return err
	}

	return nil
}
