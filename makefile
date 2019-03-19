export GO111MODULE=on
BINARY_NAME=aviasales

all: deps build
install:
	go install cmd/aviasales/aviasales.go
build:
	go build cmd/aviasales/aviasales.go
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go build -v ./...
upgrade:
	go get -u
vendor:
	go mod vendor
