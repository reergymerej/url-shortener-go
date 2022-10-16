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
		got := convert(idProvider, input)
		if input == got {
			t.Errorf("got %v", got)
		}
	})

	t.Run("should not return the same code", func(t *testing.T) {
		idProvider := &StatefulIdProvider{}
		input := "hello"
		got1 := convert(idProvider, input)
		got2 := convert(idProvider, input)
		if got1 == got2 {
			t.Errorf("got %s, %s", got1, got2)
		}
	})
}

type MockValueProvider struct {
	value string
}

func (m *MockValueProvider) GetValue(code string) string {
	return m.value
}

func TestLookup(t *testing.T) {
	t.Run("should return original string", func(t *testing.T) {
		mockValue := "good boy"
		codeValueProvider := &MockValueProvider{
			value: mockValue,
		}
		code := "a sweet code"
		got := lookup(codeValueProvider, code)
		want := mockValue
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}

func TestCodeValueProvider(t *testing.T) {
	t.Run("GetValue should return the correct value", func(t *testing.T) {
		code := "some-code"
		value := "some sweet value"
		valueMap := make(map[string]string)
		valueMap[code] = value
		codeValueProvider := &CodeValueProvider{
			valueMap: valueMap,
		}
		got := codeValueProvider.GetValue(code)
		want := value
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}
