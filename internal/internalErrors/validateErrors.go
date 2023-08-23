package internalErrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidatorStruct(obj interface{}) error {
	validation := validator.New()

	err := validation.Struct(obj)

	if err != nil {
		var validationErr validator.ValidationErrors
		errors.As(err, &validationErr)

		/*
			IDE`S recommend use errors.as

			``
		*/

		valid := validationErr[0]

		switch valid.Tag() {
		case "required":
			return errors.New(valid.StructField() + " is required")
		case "max":
			return errors.New(valid.StructField() + " is required with max " + valid.Param())
		case "min":
			return errors.New(valid.StructField() + " is required with min " + valid.Param())
		case "email":
			return errors.New(valid.StructField() + " is required")
		}
	}

	return nil
}
