VERSION?=0.4.4

all: build
build:
	CGO_ENABLED=0 go build -ldflags "-X main.Version=$(VERSION)" -v .

docker:
	docker build -t bbq:0.3 .

