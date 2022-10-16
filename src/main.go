package main

import "fmt"

type IdProvider interface {
	GetId() string
}

type ValueProvider interface {
	GetValue(code string) string
}

type StatefulIdProvider struct {
	Count int
}

func (s *StatefulIdProvider) GetId() string {
	s.Count++
	return fmt.Sprintf("id-%v", s.Count)
}

type CodeValueProvider struct {
	valueMap map[string]string
}

func (c *CodeValueProvider) GetValue(code string) string {
	return c.valueMap[code]
}

func convert(idProvider IdProvider, input string) string {
	return idProvider.GetId()
}

func lookup(valueProvider ValueProvider, code string) string {
	return valueProvider.GetValue(code)
}

func main() {
	// TODO
}
