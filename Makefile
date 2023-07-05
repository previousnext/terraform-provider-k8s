#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/terraform-provider-k8s

# Builds the project
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/terraform-provider-k8s_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' $(PROJECT)

# Run all lint checking with exit codes for CI
lint:
	golangci-lint run --disable=errcheck --timeout=5m

.PHONY: build lint
