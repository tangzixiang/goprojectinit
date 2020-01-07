.PHONY : all fmt test

all : fmt test

fmt :
	gofmt -l -e -d -w internal/ cmd/ pkg/

test :
	go test -v ./...