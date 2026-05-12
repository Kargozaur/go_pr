package validator_test

import (
	"ecommerce/user-service/internal/util/validator"
	"ecommerce/user-service/schemas"
	"strings"
	"testing"
)

func TestValidateSchema_Valid(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "Secure!1pass",
	}
	if err := v.ValidateSchema(schema); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidateSchema_InvalidEmail(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "not-an-email",
		Password: "Secure!1pass",
	}
	if err := v.ValidateSchema(schema); err == nil {
		t.Error("expected error for invalid email, got nil")
	}
}

func TestValidateSchema_InvalidPassword(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "weak",
	}
	if err := v.ValidateSchema(schema); err == nil {
		t.Error("expected error for invalid password, got nil")
	}
}

func TestValidateSchema_BothInvalid(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "not-an-email",
		Password: "weak",
	}
	err := v.ValidateSchema(schema)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "8") {
		t.Errorf("expected password error in message, got: %s", err.Error())
	}
}

func TestVerifyEmail_Valid(t *testing.T) {
	v := validator.NewValidator()
	emails := []string{
		"user@example.com",
		"user+tag@example.com",
		"user@mail.example.com",
	}
	for _, email := range emails {
		schema := schemas.UserDefaultSchema{Email: email, Password: "Secure!1pass"}
		if err := v.ValidateSchema(schema); err != nil {
			t.Errorf("email %q: expected valid, got: %v", email, err)
		}
	}
}

func TestVerifyEmail_Invalid(t *testing.T) {
	v := validator.NewValidator()
	emails := []string{
		"not-an-email",
		"user@",
		"",
		"@",
	}
	for _, email := range emails {
		schema := schemas.UserDefaultSchema{Email: email, Password: "Secure!1pass"}
		if err := v.ValidateSchema(schema); err == nil {
			t.Errorf("email %q: expected error, got nil", email)
		}
	}
}

func TestValidatePassword_TooShort(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "Abc!1",
	}
	err := v.ValidateSchema(schema)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "8") {
		t.Errorf("expected '8' in error message, got: %v", err)
	}
}

func TestValidatePassword_NoDigit(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "Secure!pass",
	}
	err := v.ValidateSchema(schema)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "number") {
		t.Errorf("expected 'number' in error, got: %v", err)
	}
}

func TestValidatePassword_NoSpecial(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "Secure1pass",
	}
	if err := v.ValidateSchema(schema); err == nil {
		t.Error("expected error for missing special char, got nil")
	}
}

func TestValidatePassword_NoUppercase(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "secure!1pass",
	}
	err := v.ValidateSchema(schema)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "upper") {
		t.Errorf("expected 'upper' in error, got: %v", err)
	}
}

func TestValidatePassword_MultipleErrors(t *testing.T) {
	v := validator.NewValidator()
	schema := schemas.UserDefaultSchema{
		Email:    "user@example.com",
		Password: "securepass",
	}
	err := v.ValidateSchema(schema)
	if err == nil {
		t.Fatal("expected multiple errors, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "number") {
		t.Errorf("expected 'number' in error, got: %s", msg)
	}
	if !strings.Contains(msg, "upper") {
		t.Errorf("expected 'upper' in error, got: %s", msg)
	}
}

func TestValidateSpecial_AcceptedChars(t *testing.T) {
	v := validator.NewValidator()
	specials := []rune{'!', '#', '$', '%', '&', '*', '+', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '^', '_', '`', '|', '~'}
	for _, ch := range specials {
		password := "Password1" + string(ch)
		schema := schemas.UserDefaultSchema{Email: "user@example.com", Password: password}
		if err := v.ValidateSchema(schema); err != nil {
			t.Errorf("char %q should be accepted as special, got: %v", ch, err)
		}
	}
}

func TestValidateSpecial_RejectedChars(t *testing.T) {
	v := validator.NewValidator()
	rejected := []rune{'(', ')', '[', ']', '{', '}', '"', '\'', '\\', ','}
	for _, ch := range rejected {
		password := "Password1" + string(ch)
		schema := schemas.UserDefaultSchema{Email: "user@example.com", Password: password}
		if err := v.ValidateSchema(schema); err == nil {
			t.Errorf("char %q should NOT be accepted as special", ch)
		}
	}
}
