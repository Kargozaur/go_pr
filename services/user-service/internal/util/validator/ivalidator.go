package validator

import "ecommerce/user-service/schemas"

type IValidator interface {
	ValidateSchema(schemas.UserDefaultSchema) error
}
