#! /bin/bash
all: test vet lint

test:
	go test -v ./...

lint:
	golint -set_exit_status ./...

coverage:
	go list ./... | xargs -n1 -I {} -P 4 go run scripts/test_with_stripe_mock/main.go -covermode=count -coverprofile=../../../{}/profile.coverprofile {}

vet:
	go vet -v .