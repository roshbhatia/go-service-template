.PHONY: check-go-version clean run docker-build docker-run 
go-version := 1.15.6
shell-go-version := $(shell go version | cut -d" " -f3 | sed -r 's/^.{2}//')
go-os := $(shell  uname | awk '{print tolower($$0)}')
binary-exists := $(shell echo $$(test -f bin/echo-service)$$?)
semver := $(shell cat VERSION)
service-port := 8080
local-port := 8080

check-go-version:
ifneq ($(strip $(go-version)),$(strip $(shell-go-version)))
	$(error error: local go version is not $(go-version))
endif

clean:
	rm -rf $(CURDIR)/bin

run: check-go-version
	go run cmd/main/main.go

build: check-go-version
	mkdir -p $(CURDIR)/bin
	env GOOS=$(go-os) GOARCH=amd64 go build -o $(CURDIR)/bin/echo-service cmd/main/main.go 

docker-build: check-go-version 
# Test for binary and recompile if OS is not linux
	if [ '$(go-os)' != 'linux' ] || [ $(binary-exists) = 1 ]; then make build go-os=linux; fi
	docker build -f ./deploy/docker/Dockerfile -t echo-service:$(semver)  .

docker-run:
	docker run -it echo-service:$(semver) -p $(local-port):$(service-port)