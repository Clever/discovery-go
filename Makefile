include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: test $(PKGS)
SHELL := /bin/bash
PKGS = $(shell go list ./...)
$(eval $(call golang-version-check,1.10))

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	@go get -d -t $@
	$(call golang-test-all-strict,$@)


install_deps: golang-dep-vendor-deps
	$(call golang-dep-vendor)