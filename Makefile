include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: test $(PKGS)
SHELL := /bin/bash
PKGS = $(shell go list ./... | grep -v /vendor)
$(eval $(call golang-version-check,1.16))

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	@go get -d -t $@
	$(call golang-test-all-strict,$@)

install_deps:
	go mod vendor
