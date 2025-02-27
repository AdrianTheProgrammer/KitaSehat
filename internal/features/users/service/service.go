package service

import (
	"KitaSehat_Backend/internal/features/users"
	"KitaSehat_Backend/internal/helper"
	"KitaSehat_Backend/internal/utils"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService struct {
	qry      users.UQuery
	pu       utils.PasswordUtilityInterface
	tu       utils.TokenUtilityInterface
	cu       utils.CloudinaryUtilityInterface
	validate *validator.Validate
}

func NewUserService(q users.UQuery, p utils.PasswordUtilityInterface, t utils.TokenUtilityInterface, c utils.CloudinaryUtilityInterface) users.UService {
	return &UserService{
		qry:      q,
		pu:       p,
		tu:       t,
		cu:       c,
		validate: validator.New(),
	}
}

func (us *UserService) Login(email string, password string) (int, string, string) {
	// Validate Login Request
	err := us.validate.Struct(&users.LoginValidate{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println("login data validation failed.", err)
		return 400, "login data validation failed", ""
	}

	// Find User by Email in Database
	result, err := us.qry.Login(email)
	if err == gorm.ErrRecordNotFound {
		log.Printf("user with email %v not found. %v", email, err)
		return 404, fmt.Sprintf("user with email %v not found.", email), ""
	} else if err != nil {
		log.Println("unexpected database error.", err)
		return 500, "unexpected database error", ""
	}

	// Compare Password
	err = us.pu.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Printf("wrong password for user %v. %v", result.Username, err)
		return 401, fmt.Sprintf("wrong password for user %v", result.Username), ""
	}

	// Generate JSON Web Token
	token, err := us.tu.GenerateToken(result)
	if err != nil {
		log.Printf("error generating jwt for login user %v. %v", result.Username, err)
		return 500, fmt.Sprintf("error generating jwt for login user %v", result.Username), ""
	}

	return 200, "login success", token
}

func (us *UserService) Register(input users.User) (int, string) {
	// Validate Register Request
	err := us.validate.Struct(&users.RegisterValidate{
		Username:    input.Username,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	})
	if err != nil {
		log.Println("register data validation failed.", err)
		return 400, "register data validation failed"
	}

	// Check if Email already existed
	if us.qry.IsEmailRegistered(input.Email) {
		log.Println("register failed. email address already registered.")
		return 409, "register failed. email address already registered"
	}

	// Generate Hashed Password
	hashedPass, err := us.pu.GeneratePassword(input.Password)
	input.Password = string(hashedPass)
	if err != nil {
		log.Println("error hashing password.", err)
		return 500, "error hashing password"
	}

	// Insert User Data Into Database
	err = us.qry.Register(input)
	if err != nil {
		log.Println("error registering user data into database.", err)
		return 500, "error registering user data into database"
	}

	return 201, "register success"
}

func (us *UserService) GetAllUsers(currentPage int) (int, string, []users.User, int) {
	// Get All User Data From Database

	limit := helper.DefaultItemsPerPage
	offset := limit * (currentPage - 1)

	result, totalItems, err := us.qry.GetAllUsers(currentPage, limit, offset)
	if err != nil {
		log.Println("error fetching user data from database.", err)
		return 500, "error fetching user data from database", []users.User{}, 0
	}

	return 200, "get all user data success", result, totalItems
}

func (us *UserService) GetUser(id int) (int, string, users.User) {
	// Get User Data By ID
	result, err := us.qry.GetUser(id)
	if err == gorm.ErrRecordNotFound {
		log.Printf("user id %v not found. %v", id, err)
		return 404, fmt.Sprintf("user id %v not found", id), users.User{}
	} else if err != nil {
		log.Println("unexpected database error.", err)
		return 500, "unexpected database error", users.User{}
	}

	return 200, "get user data success", result
}

func (us *UserService) UpdateUser(c echo.Context, id int, input users.User) (int, string) {
	// Check if operation performed by the users themselves or the Admin
	loginData := us.tu.DecodeToken(c.Get("user").(*jwt.Token))
	if id != loginData.UserID {
		log.Println("access denied. mismatched user id between request body and jwt.")
		return 401, "access denied. mismatched user id between request body and jwt"
	} else if id != loginData.UserID && loginData.AccessLevel != "developer" {
		log.Println("access denied. developer access required.")
		return 403, "access denied. developer access required"
	}

	// Validate Update Request
	err := us.validate.Struct(&users.UpdateUserValidate{
		Username:    input.Username,
		PhoneNumber: input.PhoneNumber,
	})
	if err != nil {
		log.Println("update user data validation failed.", err)
		return 400, "update user data validation failed"
	}

	// Upload Avatar to Cloudinary
	input.Avatar, err = us.cu.UploadAvatar(c, id)
	if err != nil {
		log.Println("error uploading avatar to cloudinary.", err)
		return 500, "error uploading avatar to cloudinary"
	}

	// Update User Data By ID
	err = us.qry.UpdateUser(id, input)
	if err == gorm.ErrRecordNotFound {
		log.Printf("user id %v not found. %v", id, err)
		return 404, fmt.Sprintf("user id %v not found", id)
	} else if err != nil {
		log.Println("unexpected database error.", err)
		return 500, "unexpected database error"
	}

	return 200, "update user data success"
}

func (us *UserService) DeleteUser(id int) (int, string) {
	// Delete User Data By ID
	err := us.qry.DeleteUser(id)
	if err == gorm.ErrRecordNotFound {
		log.Printf("user id %v not found. %v", id, err)
		return 404, fmt.Sprintf("user id %v not found.", id)
	} else if err != nil {
		log.Println("unexpected database error.", err)
		return 500, "unexpected database error"
	}

	return 200, "delete user data success"
}

func (us *UserService) ChangePassword(id int, input users.User) (int, string) {
	// Validate Change Password Request
	err := us.validate.Struct(&users.ChangePasswordValidate{
		Password: input.Password,
	})
	if err != nil {
		log.Println("change password data validation failed.", err)
		return 400, "change password data validation failed"
	}

	// Generate Hashed Password
	hashedPass, err := us.pu.GeneratePassword(input.Password)
	input.Password = string(hashedPass)
	if err != nil {
		log.Println("error hashing password.", err)
		return 500, "error hashing password"
	}

	// Change User Password By ID
	err = us.qry.ChangePassword(id, input)
	if err == gorm.ErrRecordNotFound {
		log.Printf("user id %v not found. %v", id, err)
		return 404, fmt.Sprintf("user id %v not found", id)
	} else if err != nil {
		log.Println("unexpected database error.", err)
		return 500, "unexpected database error"
	}

	return 200, "change password success"
}
