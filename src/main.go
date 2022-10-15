package main

type IdProvider interface {
	GetId() string
}

func convert(input string, idProvider IdProvider) string {
	return "TODO: IOU one real conversion" + idProvider.GetId()
}

func main() {
	// TODO
}
