SHELL := /bin/bash # Use bash syntax

# Set up variables
GO111MODULE=on

AWS_SDK_GO_VERSION="$(shell echo $(shell go list -m -f '{{.Version}}' github.com/aws/aws-sdk-go-v2/service/sagemaker))"
AWS_SDK_GO_VERSIONED_PATH="$(shell echo github.com/aws/aws-sdk-go-v2/service/sagemaker@$(AWS_SDK_GO_VERSION))"
SAGEMAKER_API_PATH="$(shell echo $(shell go env GOPATH))/pkg/mod/$(AWS_SDK_GO_VERSIONED_PATH)"
SERVICE_CONTROLLER_SRC_PATH="$(shell pwd)"

# Build ldflags
VERSION ?= "v0.0.0"
GITCOMMIT=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GO_LDFLAGS=-ldflags "-X main.version=$(VERSION) \
			-X main.buildHash=$(GITCOMMIT) \
			-X main.buildDate=$(BUILDDATE)"
GOCOVER=go tool cover

.PHONY: all test clean-mocks mocks

all: test

test: 				## Run code tests
	go test -v ./...

test-cover: | mocks				## Run code tests with resources coverage
	go test -coverpkg=./pkg/resource/... -covermode=count -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out -o coverage.html

# clean-mocks:	## Remove mocks directory
# 	rm -r test/mocks

# install-mockery:
# 	@test/scripts/install-mockery.sh

# mocks: install-mockery ## Build mocks
# 	go get $(AWS_SDK_GO_VERSIONED_PATH)
# 	@echo "building mocks for $(SAGEMAKER_API_PATH) ... "
# 	@pushd $(SAGEMAKER_API_PATH) 1>/dev/null; \
# 	$(SERVICE_CONTROLLER_SRC_PATH)/test/bin/mockery --all --dir=. --output=$(SERVICE_CONTROLLER_SRC_PATH)/test/mocks/aws-sdk-go-v2/sagemaker/ ; \
# 	popd 1>/dev/null;
# 	@echo "ok."


help:           	## Show this help.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v grep | sed -e 's/\\$$//' \
		| awk -F'[:#]' '{print $$1 = sprintf("%-30s", $$1), $$4}'
