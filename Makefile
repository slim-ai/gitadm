SHELL := /bin/bash
APP_NAME := gitadm

menu: ## prints out the menu of command options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(lastword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: menu

default: menu
.PHONY: default

build: build_mac build_linux build_windows ## build the app for all targets
.PHONY: build

build_mac: ## build the app for MacOS
	@bash $(CURDIR)/scripts/build.sh $(APP_NAME) darwin amd64
.PHONY: build_mac

build_linux: ## build the app for Linux
	@bash $(CURDIR)/scripts/build.sh $(APP_NAME) linux amd64
.PHONY: build_linux

build_windows: ## build the app for windows
	@bash $(CURDIR)/scripts/build.sh $(APP_NAME) windows amd64
.PHONY: build_windows

clean: ## clean up intermediate files
	@rm -fr bin release vendor
.PHONY: clean

fmt: ## automatically formats Go source code
	gofmt -l -w -s .
.PHONY: fmt

inspect: ## examines Go source code and reports suspicious constructs
	go vet ./...
.PHONY: inspect

