#! /bin/bash

test:
	go tet -v ./...

vet:
	go vet -v .