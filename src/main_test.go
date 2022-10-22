package main

import (
	"testing"
)

func TestSanity(t *testing.T) {
	main()
}

type IdProviderSpy struct {
	Calls []string
	id    string
}

func (i *IdProviderSpy) GetId(id string) string {
	i.Calls = append(i.Calls, id)
	return i.id
}

type ValueProviderSpy struct {
	Calls []string
}

func (v *ValueProviderSpy) GetValue(id string) string {
	v.Calls = append(v.Calls, id)
	return "dummy value"
}

type MockValueProvider struct {
	value string
}

func (m *MockValueProvider) GetValue(id string) string {
	return m.value
}

func TestCodeValueProvider(t *testing.T) {
	t.Run("GetValue should return the correct value", func(t *testing.T) {
		id := "some-id"
		value := "some sweet value"
		valueMap := make(map[string]string)
		valueMap[id] = value
		codeValueProvider := &CodeValueProvider{
			valueMap: valueMap,
		}
		got := codeValueProvider.GetValue(id)
		want := value
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}

func TestJelloValueProvider(t *testing.T) {
	t.Run("should return 'Jello'", func(t *testing.T) {
		jelloValueProvider := &JelloValueProvider{}
		got := jelloValueProvider.GetValue("xyz")
		want := "Jello"
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}

func TestConverter(t *testing.T) {
	t.Run("should convert from id to value", func(t *testing.T) {
		value := "personal"
		defaultConverter := &DefaultConverter{
			valueProvider: &MockValueProvider{
				value: "personal",
			},
		}
		id := "jesus"
		got := defaultConverter.GetValue(id)
		want := value
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})

	t.Run("should convert from value to id", func(t *testing.T) {
		id := "neveragain!"
		defaultConverter := &DefaultConverter{
			idProvider: &IdProviderSpy{
				id: id,
			},
		}
		value := "jesus"
		got := defaultConverter.GetId(value)
		want := id
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	})
}
