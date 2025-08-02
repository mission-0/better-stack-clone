package utilities

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mission-0/better-stack-backend/models"
)

var validationMessages = map[string]string{
	"required":    "%s is required",
	"email":       "%s must be a valid email address",
	"interval":    "Interval is must",
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
	return validate
}
