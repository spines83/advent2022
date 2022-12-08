package main

import "testing"

func TestPriority01(t *testing.T) {

	got := IsContained(2, 4, 6, 8)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPriority02(t *testing.T) {

	got := IsContained(2, 3, 4, 5)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPriority03(t *testing.T) {

	got := IsContained(5, 7, 7, 9)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPriority04(t *testing.T) {

	got := IsContained(2, 8, 3, 7)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPriority05(t *testing.T) {

	got := IsContained(6, 6, 4, 6)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPriority06(t *testing.T) {

	got := IsContained(2, 6, 4, 8)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAnyOverlap01(t *testing.T) {

	got := AnyOverlap(2, 4, 6, 8)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAnyOverlap02(t *testing.T) {

	got := AnyOverlap(2, 3, 4, 5)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAnyOverlap03(t *testing.T) {

	got := AnyOverlap(5, 7, 7, 9)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAnyOverlap04(t *testing.T) {

	got := AnyOverlap(2, 8, 3, 7)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAnyOverlap05(t *testing.T) {

	got := AnyOverlap(6, 6, 4, 6)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestAnyOverlap06(t *testing.T) {

	got := AnyOverlap(2, 6, 4, 8)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
