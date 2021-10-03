package errors

import "errors"

var (
	// ErrorValidationEmptyInput occurs when no input is provided to the delivery application
	ErrorValidationEmptyInput = errors.New("no input found")

	// ErrorValidationInvalidDelivererCount occurs when there is not at least one available deliverer
	ErrorValidationInvalidDelivererCount = errors.New("invalid number of deliverers")

	// ErrorValidationInvalidInput occurs when invalid input is detected
	ErrorValidationInvalidInput = errors.New("invalid input detected")
)
