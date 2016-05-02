package main

import (
	"testing"
)

func TestPositionsGainedPointsPositive(t *testing.T) {
	r := positionsGainedPoints(2, 1)
	if r != 1 {
		t.Error("Expected 1 got", r)
	}
}

func TestPositionsGainedPointsNegative(t *testing.T) {
	r := positionsGainedPoints(1, 2)
	if r != 0 {
		t.Error("Expected 0 got", r)
	}
}

func TestOutQualifyTeamMatePointsPositive(t *testing.T) {
	r := outQualifyTeamMatePoints(1, 2)
	if r != 5 {
		t.Error("Expected 5 got", r)
	}
}

func TestOutQualifyTeamMatePointsNegative(t *testing.T) {
	r := outQualifyTeamMatePoints(2, 1)
	if r != 0 {
		t.Error("Expected 0 got", r)
	}
}

func TestCalculatePoints(t *testing.T) {
	r := calculatePoints(2, 1, 3)
	if r != 31 {
		t.Error("Expected 31 got", r)
	}
}

func TestPositionToPointsTop10(t *testing.T) {
	r := positionToPoints(2)
	if r != 18 {
		t.Error("Expected 18 got", r)
	}
}

func TestPositionToPointsOutsideTop10(t *testing.T) {
	r := positionToPoints(12)
	if r != 0 {
		t.Error("Expected 0 got", r)
	}
}

func TestGetDriver(t *testing.T) {
	r := getDriver("HAM")
	if r.Name != "HAM" {
		t.Error("Expected HAM got", r)
	}
}
