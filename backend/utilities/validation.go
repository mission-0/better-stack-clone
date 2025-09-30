package utilities

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/mission-0/better-stack-backend/models"
)

var validationMessages = map[string]string{
	"required":    "%s is required",
	"email":       "%s must be a valid email address",
	"interval":    "%s must be a positive integer",
	"url":         "%s must be a valid URL",
	"validregion": "%s must be a valid region (Asia, Europe, North America, Middle East)",
	"min":         "%s must be at least %s characters long",
}

func FormatValidationErrors(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		msg, ok := validationMessages[err.Tag()]
		if !ok {
			msg = "%s failed validation with tag %s"
		}
		fieldName := err.Field()
		if err.Tag() == "min" {
			errors = append(errors, fmt.Sprintf(msg, fieldName, err.Param()))
		} else {
			errors = append(errors, fmt.Sprintf(msg, fieldName))
		}
	}
	return errors
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("validregion", func(fl validator.FieldLevel) bool {
		region := fl.Field().Interface().(models.RegionList)
		return models.IsValidRegion(region)
	})
	// interval: must be a positive integer (> 0)
	validate.RegisterValidation("interval", func(fl validator.FieldLevel) bool {
		// handles all int kinds by using Int() which returns int64
		// If the field is not an integer kind, validation fails
		switch fl.Field().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return fl.Field().Int() > 0
		default:
			return false
		}
	})
	return validate
}
