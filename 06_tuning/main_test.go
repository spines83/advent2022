package main

import "testing"

func TestMarker01(t *testing.T) {

	got, _ := GetFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMarker02(t *testing.T) {

	got, _ := GetFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMarker03(t *testing.T) {

	got, _ := GetFirstMarker("nppdvjthqldpwncqszvftbrmjlhg", 4)
	want := 6

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMarker04(t *testing.T) {

	got, _ := GetFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMarker05(t *testing.T) {

	got, _ := GetFirstMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPacket01(t *testing.T) {

	got, _ := GetFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
	want := 19

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPacket02(t *testing.T) {

	got, _ := GetFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14)
	want := 23

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPacket03(t *testing.T) {

	got, _ := GetFirstMarker("nppdvjthqldpwncqszvftbrmjlhg", 14)
	want := 23

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPacket04(t *testing.T) {

	got, _ := GetFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14)
	want := 29

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPacket05(t *testing.T) {

	got, _ := GetFirstMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14)
	want := 26

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
