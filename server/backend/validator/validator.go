package validator

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateName(name string) error {
	return validation.Validate(name,
		validation.Required,
		validation.Match(regexp.MustCompile("^[a-zA-Z ]+$")).Error("name must not contain numbers or special characters"),
	)
}
