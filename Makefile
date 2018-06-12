install:
	go install -v

build:
	go build -v ./...

deps: dev-deps
	go get -u github.com/nats-io/go-nats

dev-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/smartystreets/goconvey/convey

test:
	go test -v ./...

lint:
	golint ./...
	go vet ./...
