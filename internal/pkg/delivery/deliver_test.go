package delivery_test

import (
	appErrors "pizza-delivery/internal/errors"
	"pizza-delivery/internal/pkg/delivery"
	"testing"
)

func TestSuccessfulDeliveryOneInstructionOneDeliverer(t *testing.T) {
	validInput := ">"
	validDelivererCount := 1
	housesVisited, err := delivery.Deliver(validInput, validDelivererCount)

	if err != nil {
		t.Error("an unexpected error occurred", err)
	}

	expectedHousesVisitedCount := 2
	if housesVisited != expectedHousesVisitedCount {
		t.Errorf("expected to visit %d but visited %d instead", expectedHousesVisitedCount, housesVisited)
	}
}

func TestSuccessfulDeliveryAroundTheBlockOneDeliverer(t *testing.T) {
	validInput := "^>v<" //nolint:goconst
	validDelivererCount := 1
	housesVisited, err := delivery.Deliver(validInput, validDelivererCount)

	if err != nil {
		t.Error("an unexpected error occurred", err)
	}

	expectedHousesVisitedCount := 4
	if housesVisited != expectedHousesVisitedCount {
		t.Errorf("expected to visit %d but visited %d instead", expectedHousesVisitedCount, housesVisited)
	}
}

func TestSuccessfulDeliveryBackAndForthOneDeliverer(t *testing.T) {
	validInput := "^v^v^v^v^v" //nolint:goconst
	validDelivererCount := 1
	housesVisited, err := delivery.Deliver(validInput, validDelivererCount)

	if err != nil {
		t.Error("an unexpected error occurred", err)
	}

	expectedHousesVisitedCount := 2
	if housesVisited != expectedHousesVisitedCount {
		t.Errorf("expected to visit %d but visited %d instead", expectedHousesVisitedCount, housesVisited)
	}
}

func TestSuccessfulDeliveryOneInstructionTwoDeliverers(t *testing.T) {
	validInput := "^v"
	validDelivererCount := 2
	housesVisited, err := delivery.Deliver(validInput, validDelivererCount)

	if err != nil {
		t.Error("an unexpected error occurred", err)
	}

	expectedHousesVisitedCount := 3
	if housesVisited != expectedHousesVisitedCount {
		t.Errorf("expected to visit %d but visited %d instead", expectedHousesVisitedCount, housesVisited)
	}
}

func TestSuccessfulDeliveryAroundTheBlockTwoDeliverers(t *testing.T) {
	validInput := "^>v<"
	validDelivererCount := 2
	housesVisited, err := delivery.Deliver(validInput, validDelivererCount)

	if err != nil {
		t.Error("an unexpected error occurred", err)
	}

	expectedHousesVisitedCount := 3
	if housesVisited != expectedHousesVisitedCount {
		t.Errorf("expected to visit %d but visited %d instead", expectedHousesVisitedCount, housesVisited)
	}
}

func TestSuccessfulDeliveryBackAndForthTwoDeliverers(t *testing.T) {
	validInput := "^v^v^v^v^v"
	validDelivererCount := 2
	housesVisited, err := delivery.Deliver(validInput, validDelivererCount)

	if err != nil {
		t.Error("an unexpected error occurred", err)
	}

	expectedHousesVisitedCount := 11
	if housesVisited != expectedHousesVisitedCount {
		t.Errorf("expected to visit %d but visited %d instead", expectedHousesVisitedCount, housesVisited)
	}
}

func TestEmptyInputFailure(t *testing.T) {
	invalidInput := ""
	validDelivererCount := 1
	_, err := delivery.Deliver(invalidInput, validDelivererCount)

	if err != appErrors.ErrorValidationEmptyInput {
		t.Error("expected an error based on empty input to the delivery application")
	}
}

func TestInvalidDelivererCountFailure(t *testing.T) {
	validInput := ">"
	invalidDelivererCount := -1
	_, err := delivery.Deliver(validInput, invalidDelivererCount)

	if err != appErrors.ErrorValidationInvalidDelivererCount {
		t.Error("expected an error based on an invalid number of deliverers")
	}
}
