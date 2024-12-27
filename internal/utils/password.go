package utils

import "golang.org/x/crypto/bcrypt"

type PasswordUtilityInterface interface {
	GeneratePassword(string) ([]byte, error)
	ComparePassword([]byte, []byte) error
}

type PasswordUtility struct{}

func NewPasswordUtility() PasswordUtilityInterface {
	return &PasswordUtility{}
}

func (pu *PasswordUtility) GeneratePassword(input string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pu *PasswordUtility) ComparePassword(current, input []byte) error {
	return bcrypt.CompareHashAndPassword(current, input)
}
