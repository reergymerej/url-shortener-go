package main

import "fmt"

type IdProvider interface {
	GetId() string
}

type StatefulIdProvider struct {
	Count int
}

func (s *StatefulIdProvider) GetId() string {
	s.Count++
	return fmt.Sprintf("id-%v", s.Count)
}

func convert(input string, idProvider IdProvider) string {
	return idProvider.GetId()
}

func main() {
	// TODO
}
