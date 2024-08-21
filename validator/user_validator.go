package validator

import (
	"rest-api-udemy/model"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct {}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (tv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.Length(1, 30).Error("limited max 30 char"),
			is.Email.Error("email is invalid"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.Length(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}
