package validation_test

import (
	appErrors "pizza-delivery/internal/errors"
	"pizza-delivery/internal/pkg/validation"
	"testing"
)

func TestEmptyInputFailure(t *testing.T) {
	invalidInput := ""
	validDelivererCount := 1
	err := validation.Validate(invalidInput, validDelivererCount)

	if err != appErrors.ErrorValidationEmptyInput {
		t.Error("expected a validation failure base on empty input")
	}
}

func TestInvalidDelivererCountFailure(t *testing.T) {
	validInput := ">"
	invalidDelivererCount := 0
	err := validation.Validate(validInput, invalidDelivererCount)

	if err != appErrors.ErrorValidationInvalidDelivererCount {
		t.Error("expected a validation failure base an invalid number of deliverers")
	}
}

func TestInvalidInputFailure(t *testing.T) {
	validInput := "><^^%"
	invalidDelivererCount := 1
	err := validation.Validate(validInput, invalidDelivererCount)

	if err != appErrors.ErrorValidationInvalidInput {
		t.Error("expected a validation failure base on invalid input")
	}
}

func TestValidationSuccess(t *testing.T) {
	validInput := "><^^"
	invalidDelivererCount := 1
	err := validation.Validate(validInput, invalidDelivererCount)

	if err != nil {
		t.Error("validation failed unexpectedly", err)
	}
}
