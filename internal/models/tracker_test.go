package models_test

import (
	"pizza-delivery/internal/models"
	"testing"
)

func TestTrackerDefault(t *testing.T) {
	tracker := models.NewTracker()
	uniqueLocations := tracker.GetUniqueLocations()

	expectedCount := 0
	if len(uniqueLocations) != expectedCount {
		t.Errorf("expected %d unique locations but found %d", expectedCount, len(uniqueLocations))
	}
}

func TestTrackerWithOneLocation(t *testing.T) {
	location := models.NewLocation(0, 0)
	tracker := models.NewTracker()
	tracker.AddLocation(location)
	uniqueLocations := tracker.GetUniqueLocations()

	expectedCount := 1
	if len(uniqueLocations) != expectedCount {
		t.Errorf("expected %d unique location but found %d", expectedCount, len(uniqueLocations))
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
	uniqueLocations := tracker.GetUniqueLocations()

	expectedCount := 3
	if len(uniqueLocations) != expectedCount {
		t.Errorf("expected %d unique locations but found %d", expectedCount, len(uniqueLocations))
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
	uniqueLocations := tracker.GetUniqueLocations()

	expectedCount := 2
	if len(uniqueLocations) != expectedCount {
		t.Errorf("expected %d unique locations but found %d", expectedCount, len(uniqueLocations))
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
	uniqueLocations := tracker.GetUniqueLocations()

	if len(uniqueLocations) != 1 {
		t.Errorf("expected 1 unique location but found %d", len(uniqueLocations))
	}
}
