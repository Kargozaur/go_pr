package validator

import (
	"ecommerce/user-service/internal/schemas"
)

type IValidator interface {
	ValidateSchema(schemas.UserDefaultSchema) error
	VerifyEmail(string) error
	ValidatePassword(string) error
}
