package validator

import (
	"ecommerce/user-service/internal/schemas"
	"errors"
	"net/mail"
	"unicode"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateSchema(schema schemas.UserDefaultSchema) error {
	result := make([]error, 0, 2)
	if err := v.VerifyEmail(schema.Email); err != nil {
		result = append(result, err)
	}
	if err := v.ValidatePassword(schema.Password); err != nil {
		result = append(result, err)
	}
	if len(result) != 0 {
		return errors.Join(result...)
	}
	return nil
}

func (v *Validator) VerifyEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}

func (v *Validator) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters.") // early exit before for loop
	}
	result := make([]error, 0, 3)
	var isSpecial, isDigit, isUpper bool
	for _, char := range password {
		if unicode.IsDigit(char) {
			isDigit = true
			continue
		}
		if unicode.IsUpper(char) {
			isUpper = true
			continue
		}
		if v.validateSpecial(char) {
			isSpecial = true
			continue
		}
	}
	if !isDigit {
		result = append(result, errors.New("Password must contain at least one number."))
	}
	if !isSpecial {
		result = append(result, errors.New("Password must contain at least one special symbol."))
	}
	if !isUpper {
		result = append(result, errors.New("Password must contain at least one upper case character."))
	}
	return errors.Join(result...)
}

func (v *Validator) validateSpecial(char rune) bool {
	switch char {
	case '!', '#', '$', '%', '&', '*', '+', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '^', '_', '`', '|', '~':
		return true
	default:
		return false
	}
}
