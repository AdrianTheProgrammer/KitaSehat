package service_test

import (
	"KitaSehat_Backend/internal/features/users"
	"KitaSehat_Backend/internal/features/users/service"
	"KitaSehat_Backend/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func InitMocks(t *testing.T) (*mocks.UQuery, *mocks.PasswordUtilityInterface, *mocks.TokenUtilityInterface, *mocks.CloudinaryUtilityInterface, users.UService) {
	qry := mocks.NewUQuery(t)
	pu := mocks.NewPasswordUtilityInterface(t)
	tu := mocks.NewTokenUtilityInterface(t)
	cu := mocks.NewCloudinaryUtilityInterface(t)
	srv := service.NewUserService(qry, pu, tu, cu)

	return qry, pu, tu, cu, srv
}

func TestLogin(t *testing.T) {
	qry, pu, tu, _, srv := InitMocks(t)

	user := users.User{
		Username: "Admin",
		Email:    "admin@kitasehat.com",
		Password: "123456",
	}

	t.Run("Login Validation Failed", func(t *testing.T) {
		code, msg, token := srv.Login("", "")

		assert.Equal(t, 400, code)
		assert.Equal(t, "login data validation failed", msg)
		assert.Equal(t, "", token)
	})

	t.Run("Email Not Found", func(t *testing.T) {
		qry.On("Login", user.Email).Return(users.User{}, gorm.ErrRecordNotFound).Once()

		code, msg, token := srv.Login(user.Email, user.Password)

		assert.Equal(t, 404, code)
		assert.Equal(t, fmt.Sprintf("user with email %v not found.", user.Email), msg)
		assert.Equal(t, "", token)
	})

	t.Run("Database Error", func(t *testing.T) {
		qry.On("Login", user.Email).Return(users.User{}, assert.AnError).Once()

		code, msg, token := srv.Login(user.Email, user.Password)

		assert.Equal(t, 500, code)
		assert.Equal(t, "unexpected database error", msg)
		assert.Equal(t, "", token)
	})

	t.Run("Wrong Password", func(t *testing.T) {
		qry.On("Login", user.Email).Return(user, nil).Once()
		pu.On("ComparePassword", []byte(user.Password), []byte(user.Password)).Return(assert.AnError).Once()

		code, msg, token := srv.Login(user.Email, user.Password)

		assert.Equal(t, 401, code)
		assert.Equal(t, fmt.Sprintf("wrong password for user %v", user.Username), msg)
		assert.Equal(t, "", token)
	})

	t.Run("JWT Error", func(t *testing.T) {
		qry.On("Login", user.Email).Return(user, nil).Once()
		pu.On("ComparePassword", []byte(user.Password), []byte(user.Password)).Return(nil).Once()
		tu.On("GenerateToken", user).Return("", assert.AnError).Once()

		code, msg, token := srv.Login(user.Email, user.Password)

		assert.Equal(t, 500, code)
		assert.Equal(t, fmt.Sprintf("error generating jwt for login user %v", user.Username), msg)
		assert.Equal(t, "", token)
	})

	t.Run("Login Success", func(t *testing.T) {
		qry.On("Login", user.Email).Return(user, nil).Once()
		pu.On("ComparePassword", []byte(user.Password), []byte(user.Password)).Return(nil).Once()
		tu.On("GenerateToken", user).Return("token", nil).Once()

		code, msg, token := srv.Login(user.Email, user.Password)

		assert.Equal(t, 200, code)
		assert.Equal(t, "login success", msg)
		assert.Equal(t, "token", token)
	})
}

func TestRegister(t *testing.T) {
	qry, pu, _, _, srv := InitMocks(t)

	user := users.User{
		Username:    "Admin",
		Email:       "admin@kitasehat.com",
		Password:    "123456",
		PhoneNumber: "+628987654321",
	}

	t.Run("Email Already Registered", func(t *testing.T) {
		qry.On("IsEmailRegistered", user.Email).Return(true).Once()

		code, msg := srv.Register(user)

		assert.Equal(t, 409, code)
		assert.Equal(t, "register failed. email address already registered", msg)
	})

	t.Run("Generate Password Error", func(t *testing.T) {
		qry.On("IsEmailRegistered", user.Email).Return(false).Once()
		pu.On("GeneratePassword", user.Password).Return([]byte(""), assert.AnError).Once()

		code, msg := srv.Register(user)

		assert.Equal(t, 500, code)
		assert.Equal(t, "error hashing password", msg)
	})

	t.Run("Database Error", func(t *testing.T) {
		qry.On("IsEmailRegistered", user.Email).Return(false).Once()
		pu.On("GeneratePassword", user.Password).Return([]byte(user.Password), nil).Once()
		qry.On("Register", user).Return(assert.AnError).Once()

		code, msg := srv.Register(user)

		assert.Equal(t, 500, code)
		assert.Equal(t, "error registering user data into database", msg)
	})

	t.Run("Register Success", func(t *testing.T) {
		qry.On("IsEmailRegistered", user.Email).Return(false).Once()
		pu.On("GeneratePassword", user.Password).Return([]byte(user.Password), nil).Once()
		qry.On("Register", user).Return(nil).Once()

		code, msg := srv.Register(user)

		assert.Equal(t, 201, code)
		assert.Equal(t, "register success", msg)
	})
}

func TestGetAllUsers(t *testing.T) {
	qry, _, _, _, srv := InitMocks(t)

	currentPage := 1
	limit := 20
	offset := 0

	t.Run("GetAllUsers Success", func(t *testing.T) {
		qryResult := []users.User{{UserID: 1}}
		qry.On("GetAllUsers", currentPage, limit, offset).Return(qryResult, 1, nil).Once()

		code, msg, result, totalItems := srv.GetAllUsers(currentPage)

		assert.Equal(t, 200, code)
		assert.Equal(t, "get all user data success", msg)
		assert.Equal(t, qryResult, result)
		assert.Equal(t, 1, totalItems)
	})

	t.Run("Database Error", func(t *testing.T) {
		qry.On("GetAllUsers", currentPage, limit, offset).Return([]users.User{}, 0, assert.AnError)

		code, msg, result, totalItems := srv.GetAllUsers(currentPage)

		assert.Equal(t, 500, code)
		assert.Equal(t, "error fetching user data from database", msg)
		assert.Equal(t, []users.User{}, result)
		assert.Equal(t, 0, totalItems)
	})
}

func TestGetUser(t *testing.T) {
	qry, _, _, _, srv := InitMocks(t)

	id := 1

	t.Run("User Not Found", func(t *testing.T) {
		qry.On("GetUser", id).Return(users.User{}, gorm.ErrRecordNotFound).Once()

		code, msg, result := srv.GetUser(id)

		assert.Equal(t, 404, code)
		assert.Equal(t, fmt.Sprintf("user id %v not found", id), msg)
		assert.Equal(t, users.User{}, result)
	})

	t.Run("Database Error", func(t *testing.T) {
		qry.On("GetUser", id).Return(users.User{}, assert.AnError).Once()

		code, msg, result := srv.GetUser(id)

		assert.Equal(t, 500, code)
		assert.Equal(t, "unexpected database error", msg)
		assert.Equal(t, users.User{}, result)
	})

	t.Run("GetUser Success", func(t *testing.T) {
		qry.On("GetUser", id).Return(users.User{}, nil).Once()

		code, msg, result := srv.GetUser(id)

		assert.Equal(t, 200, code)
		assert.Equal(t, "get user data success", msg)
		assert.Equal(t, users.User{}, result)
	})
}

// func TestUpdateUser(t *testing.T) {
// 	qry, _, _, cu, srv := InitMocks(t)

// 	var id int = 1
// 	var input = users.User{UserID: 1}
// 	var context jwt.Token
// 	var pointerContext = &context

// 	t.Run("Cloudinary Error", func(t *testing.T) {
// 		cu.On("UploadAvatar", c, id).Return("", assert.AnError).Once()

// 		code, msg := srv.UpdateUser(c, id, input)

// 		assert.Equal(t, 500, code)
// 		assert.Equal(t, "error uploading avatar to cloudinary", msg)
// 	})

// 	t.Run("User ID Not Found", func(t *testing.T) {
// 		cu.On("UploadAvatar", ctx, id).Return("avatar/url", nil).Once()
// 		qry.On("UpdateUser", id, input).Return(gorm.ErrRecordNotFound).Once()

// 		code, msg := srv.UpdateUser(ctx, id, input)

// 		assert.Equal(t, 404, code)
// 		assert.Equal(t, fmt.Sprintf("user id %v not found", id), msg)
// 	})
// }
