package validation

import (
	appErrors "pizza-delivery/internal/errors"
	"regexp"
)

var validDirectionsRegex = regexp.MustCompile(`^[\^v<>]+$`)

// Validate validates string input
func Validate(input string, delivererCount int) error {
	err := validateInput(input)
	if err != nil {
		return err
	}

	err = validateDeliverCount(delivererCount)
	if err != nil {
		return err
	}

	return nil
}

func validateInput(input string) error {
	if input == "" {
		return appErrors.ErrorValidationEmptyInput
	}

	if !validDirectionsRegex.MatchString(input) {
		return appErrors.ErrorValidationInvalidInput
	}

	return nil
}

func validateDeliverCount(count int) error {
	if count < 1 {
		return appErrors.ErrorValidationInvalidDelivererCount
	}

	return nil
}
