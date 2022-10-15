package main

import (
	"fmt"
	"testing"
)

func TestSanity(t *testing.T) {
	main()
}

func TestConvert(t *testing.T) {
	t.Run("should return a string", func(t *testing.T) {
		input := "hello"
		idProvider := func(input string) string {
			return "fake-id"
		}
		got := convert(input, idProvider)
		want := "TODO: IOU one real conversion" + idProvider(input)
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("should be different than input", func(t *testing.T) {
		input := "hello"
		idProvider := func(input string) string {
			return "fake-id"
		}
		got := convert(input, idProvider)
		if input == got {
			t.Errorf("got %v", got)
		}
	})

	t.Run("should not return the same code", func(t *testing.T) {
		state := 0
		idProvider := func(input string) string {
			state++
			return fmt.Sprintf("id-%v", state)
		}
		input := "hello"
		got1 := convert(input, idProvider)
		got2 := convert(input, idProvider)
		if got1 == got2 {
			t.Errorf("got %s, %s", got1, got2)
		}
	})
}
