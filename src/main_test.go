package main

import (
	"testing"
)

func TestSanity(t *testing.T) {
	main()
}

func TestConvert(t *testing.T) {
	t.Run("should be different than input", func(t *testing.T) {
		input := "hello"
		idProvider := &StatefulIdProvider{}
		got := convert(input, idProvider)
		if input == got {
			t.Errorf("got %v", got)
		}
	})

	t.Run("should not return the same code", func(t *testing.T) {
		idProvider := &StatefulIdProvider{}
		input := "hello"
		got1 := convert(input, idProvider)
		got2 := convert(input, idProvider)
		if got1 == got2 {
			t.Errorf("got %s, %s", got1, got2)
		}
	})
}
