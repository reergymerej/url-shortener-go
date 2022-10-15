package main

import "testing"

func TestSanity(t *testing.T) {
	main()
}

func TestConvert(t *testing.T) {
	got := convert()
	want := "TODO: IOU one real conversion"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
