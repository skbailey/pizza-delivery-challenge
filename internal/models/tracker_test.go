package models_test

import (
	"pizza-delivery/internal/models"
	"testing"
)

func TestTrackerDefault(t *testing.T) {
	tracker := models.NewTracker()
	uniqueLocationCount := tracker.CountUniqueLocations()

	expectedCount := 0
	if uniqueLocationCount != expectedCount {
		t.Errorf("expected %d unique locations but found %d", expectedCount, uniqueLocationCount)
	}
}

func TestTrackerWithOneLocation(t *testing.T) {
	location := models.NewLocation(0, 0)
	tracker := models.NewTracker()
	tracker.AddLocation(location)
	uniqueLocationCount := tracker.CountUniqueLocations()

	expectedCount := 1
	if uniqueLocationCount != expectedCount {
		t.Errorf("expected %d unique location but found %d", expectedCount, uniqueLocationCount)
	}
}

func TestTrackerWithMultipleUniqueLocations(t *testing.T) {
	initial := models.NewLocation(0, 0)
	second := models.NewLocation(2, 3)
	last := models.NewLocation(1, 1)

	tracker := models.NewTracker()
	tracker.AddLocation(initial)
	tracker.AddLocation(second)
	tracker.AddLocation(last)
	uniqueLocationCount := tracker.CountUniqueLocations()

	expectedCount := 3
	if uniqueLocationCount != expectedCount {
		t.Errorf("expected %d unique locations but found %d", expectedCount, uniqueLocationCount)
	}
}

func TestTrackerWithSomeDuplicateLocations(t *testing.T) {
	initial := models.NewLocation(0, 0)
	second := models.NewLocation(2, 3)
	last := models.NewLocation(2, 3)

	tracker := models.NewTracker()
	tracker.AddLocation(initial)
	tracker.AddLocation(second)
	tracker.AddLocation(last)
	uniqueLocationCount := tracker.CountUniqueLocations()

	expectedCount := 2
	if uniqueLocationCount != expectedCount {
		t.Errorf("expected %d unique locations but found %d", expectedCount, uniqueLocationCount)
	}
}

func TestTrackerWithOnlyDuplicateLocations(t *testing.T) {
	initial := models.NewLocation(2, 3)
	second := models.NewLocation(2, 3)
	last := models.NewLocation(2, 3)

	tracker := models.NewTracker()
	tracker.AddLocation(initial)
	tracker.AddLocation(second)
	tracker.AddLocation(last)
	uniqueLocationCount := tracker.CountUniqueLocations()

	if uniqueLocationCount != 1 {
		t.Errorf("expected 1 unique location but found %d", uniqueLocationCount)
	}
}
