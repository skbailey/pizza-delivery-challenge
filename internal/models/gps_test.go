package models_test

import (
	"pizza-delivery/internal/models"
	"testing"
)

func TestMoveNorth(t *testing.T) {
	location := models.NewLocation(0, 0)
	gps := models.NewGPS(location)

	newLocation, _ := gps.Move("^")
	expectedLocation := models.NewLocation(0, 1)
	if newLocation != expectedLocation {
		t.Error("expected to move north but that did not happen")
	}
}

func TestMoveSouth(t *testing.T) {
	location := models.NewLocation(0, 0)
	gps := models.NewGPS(location)

	newLocation, _ := gps.Move("v")
	expectedLocation := models.NewLocation(0, -1)
	if newLocation != expectedLocation {
		t.Error("expected to move south but that did not happen")
	}
}

func TestMoveEast(t *testing.T) {
	location := models.NewLocation(0, 0)
	gps := models.NewGPS(location)

	newLocation, _ := gps.Move(">")
	expectedLocation := models.NewLocation(1, 0)
	if newLocation != expectedLocation {
		t.Error("expected to move east but that did not happen")
	}
}

func TestMoveWest(t *testing.T) {
	location := models.NewLocation(0, 0)
	gps := models.NewGPS(location)

	newLocation, _ := gps.Move("<")
	expectedLocation := models.NewLocation(-1, 0)
	if newLocation != expectedLocation {
		t.Error("expected to move west but that did not happen")
	}
}

func TestInvalidDirection(t *testing.T) {
	location := models.NewLocation(0, 0)
	gps := models.NewGPS(location)

	_, err := gps.Move("&")
	if err == nil {
		t.Error("expected an error to be thrown")
	}
}
