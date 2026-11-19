.PHONY: build test

build:
	GOWORK=off go build -mod=readonly -o wangsu-certificate.exe .

test:
	GOWORK=off go test -mod=readonly ./...
