package entity

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID           string `gorm:"primaryKey"`
	FirstName    *string
	LastName     *string
	Gender       *string
	Role         *string
	Category     *string
	Email        *string `gorm:"unique"`
	Phone        string  `gorm:"unique"`
	Username     *string `gorm:"unique"`
	Password     *string
	Token        *string
	RefreshToken *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u User) isPhoneValid() bool {
	return len(u.Phone) == 10
}

func (u User) isEmailValid() bool {
	if u.Email == nil {
		return false
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(*u.Email)
}

func (u User) areRequiredFieldsPresent() bool {
	return u.FirstName != nil && u.LastName != nil && u.Email != nil && u.Gender != nil && u.Role != nil && u.Category != nil && u.Password != nil
}

func (u User) isFieldLengthValid(field *string, minLength int, maxLength int) bool {
	if field == nil {
		return false
	}
	length := len(*field)
	return length >= minLength && length <= maxLength
}

func (u User) IsValidUser() error {
	if !u.isPhoneValid() {
		return errors.New("phone invalid")
	}
	if !u.isEmailValid() {
		return errors.New("email invalid")
	}
	if !u.areRequiredFieldsPresent() {
		return errors.New("required fields are missing")
	}
	if !u.isFieldLengthValid(u.Username, 3, 20) {
		return errors.New("username length invalid")
	}
	if !u.isFieldLengthValid(u.Password, 8, 64) {
		return errors.New("password length invalid")
	}
	return nil
}
