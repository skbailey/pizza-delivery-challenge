package validation_test

import (
	appErrors "pizza-delivery/internal/errors"
	"pizza-delivery/internal/pkg/validation"
	"testing"
)

func TestDeliveryParamsEmptyInputFailure(t *testing.T) {
	emptyInput := ""
	validDelivererCount := 1
	err := validation.ValidateDeliveryParams(emptyInput, validDelivererCount)

	if err != appErrors.ErrorValidationEmptyInput {
		t.Error("expected a validation failure based on empty input")
	}
}

func TestDeliveryParamsInvalidDelivererCountFailure(t *testing.T) {
	validInput := ">"
	invalidDelivererCount := 0
	err := validation.ValidateDeliveryParams(validInput, invalidDelivererCount)

	if err != appErrors.ErrorValidationInvalidDelivererCount {
		t.Error("expected a validation failure based on an invalid number of deliverers")
	}
}

func TestDeliveryParamsInvalidInputFailure(t *testing.T) {
	invalidInput := "><^^%"
	validDelivererCount := 1
	err := validation.ValidateDeliveryParams(invalidInput, validDelivererCount)

	if err != appErrors.ErrorValidationInvalidInput {
		t.Error("expected a validation failure based on invalid input")
	}
}

func TestDeliveryParamsValidationSuccess(t *testing.T) {
	validInput := "><^^"
	validDelivererCount := 1
	err := validation.ValidateDeliveryParams(validInput, validDelivererCount)

	if err != nil {
		t.Error("validation failed unexpectedly", err)
	}
}

func TestReadParamsEmptyPathFailure(t *testing.T) {
	emptyPath := ""
	err := validation.ValidateReadParams(emptyPath)

	if err != appErrors.ErrorValidationEmptyInput {
		t.Error("expected a validation failure based on an empty path")
	}
}
