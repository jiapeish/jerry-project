package main

import (
	"testing"
)

func assertEq(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("\n got: %s\nwant: %s", got, want)
	}
}

func TestAdd01(t *testing.T) {
	segments := NewIntensitySegments()
	segments.Add(10, 20, 1)
	segments.Add(20, 30, 1)
	assertEq(t, segments.ToString(), "[[10,1],[30,0]]")
}

func TestAdd02(t *testing.T) {
	segments := NewIntensitySegments()
	segments.Add(10, 50, 1)
	segments.Add(20, 40, 2)
	assertEq(t, segments.ToString(), "[[10,1],[20,3],[40,1],[50,0]]")
}

func TestAdd03(t *testing.T) {
	segments := NewIntensitySegments()
	segments.Add(10, 30, 5)
	segments.Add(20, 40, -3)
	assertEq(t, segments.ToString(), "[[10,5],[20,2],[30,-3],[40,0]]")
}

func TestSet01(t *testing.T) {
	segments := NewIntensitySegments()
	segments.Add(10, 30, 1)
	segments.Add(20, 40, 1)
	segments.Set(10, 40, -2)
	assertEq(t, segments.ToString(), "[[10,-2],[40,0]]")
}

func TestSet02(t *testing.T) {
	segments := NewIntensitySegments()
	segments.Add(0, 10, 1)
	segments.Add(40, 50, 2)
	segments.Add(10, 40, 7)
	segments.Set(10, 40, 3)
	assertEq(t, segments.ToString(), "[[0,1],[10,3],[40,2],[50,0]]")
}
