package main

type IdProvider interface {
	GetId(value string) string
}

type ValueProvider interface {
	GetValue(id string) string
}

type CodeValueProvider struct {
	valueMap map[string]string
}

func (c *CodeValueProvider) GetValue(id string) string {
	return c.valueMap[id]
}

type JelloValueProvider struct{}

func (j *JelloValueProvider) GetValue(id string) string {
	return "Jello"
}

type DefaultConverter struct {
	valueProvider ValueProvider
	idProvider    IdProvider
}

func (d *DefaultConverter) GetValue(id string) string {
	return d.valueProvider.GetValue(id)
}

func (d *DefaultConverter) GetId(value string) string {
	return d.idProvider.GetId(value)
}

func main() {
	// TODO
}
