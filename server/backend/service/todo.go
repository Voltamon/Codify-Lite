package service

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TodoRequest struct {
	Description string `json:"description"`
}

func (t TodoRequest) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Description, validation.Required, validation.Length(3, 100)),
	)
}
