package main

import (
	"testing"
)

func TestPriority01(t *testing.T) {

	got := GetPriority('p')
	want := 16

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPriority02(t *testing.T) {

	got := GetPriority('L')
	want := 38

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPriority03(t *testing.T) {

	got := GetPriority('P')
	want := 42

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPriority04(t *testing.T) {

	got := GetPriority('v')
	want := 22

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPriority05(t *testing.T) {

	got := GetPriority('t')
	want := 20

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestPriority06(t *testing.T) {

	got := GetPriority('s')
	want := 19

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCommonItem01(t *testing.T) {

	input := []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}
	got, _ := FindCommonItem(input)
	want := 'p'

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCommonItem02(t *testing.T) {

	input := []string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}
	got, _ := FindCommonItem(input)
	want := 'L'

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCommonItem03(t *testing.T) {

	input := []string{"PmmdzqPrV", "vPwwTWBwg"}
	got, _ := FindCommonItem(input)
	want := 'P'

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCommonItem04(t *testing.T) {

	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}
	got, _ := FindCommonItem(input)
	want := 'r'

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
