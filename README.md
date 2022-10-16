# url-shortener-go

This is a url shortener written in Golang.

## How to Use

## How to Dev

### Test
```
cd src
go test . -cover

# see coverage
go test . -cover -v -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```
