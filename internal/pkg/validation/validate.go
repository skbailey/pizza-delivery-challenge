package validation

import "errors"

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
		return errors.New("no input found")
	}

	// TODO: Validate input with regex

	return nil
}

func validateDeliverCount(count int) error {
	if count < 1 {
		return errors.New("invalid number of deliverers")
	}

	return nil
}
