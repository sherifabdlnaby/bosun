package config

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	ens "github.com/go-playground/validator/translations/en"
	"github.com/robfig/cron/v3"
	"gopkg.in/go-playground/validator.v9"
)

// Validate
func validate(config Config) error {
	validate := validator.New()

	// Get English Errors
	uni := ut.New(en.New(), en.New())
	trans, _ := uni.GetTranslator("en")
	_ = ens.RegisterDefaultTranslations(validate, trans)

	// Validate
	err := validate.Struct(&config)

	if err != nil {
		errorLists, ok := err.(validator.ValidationErrors)
		if ok {
			return fmt.Errorf("errors validating struct config: %v", errorLists.Translate(trans))
		}

		return fmt.Errorf("errors validating struct config: %v", err.Error())
	}

	err = customValidate(config)
	if err != nil {
		return fmt.Errorf("errors validating struct config: %s", err.Error())
	}

	return nil
}

// Custom Validators
func customValidate(config Config) error {
	// Put Custom Validation Here
	// TODO Validate Regex Expressions to not contain multiple ** or ## and is a valid index pattern.

	// Validate cron schedules
	_, err := cron.ParseStandard(config.AutoIndexPattern.Schedule)
	if err != nil {
		return fmt.Errorf("cron expression not valid: %s", err.Error())
	}

	return nil
}