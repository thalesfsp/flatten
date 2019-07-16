###
# Author: Thales Pinheiro <thalesfsp@gmail.com.>
# Since: 07/2019
# Basic usage:
#  make
#  make {option}
###

ifneq (,)
	This makefile requires GNU Make.
endif

APP:=flatten

GOOS_LINUX:=linux
GOOS_MAC:=darwin
GOARCH_LINUX:=amd64
GOARCH_MAC:=amd64

BIN_FOLDER:=bin
BIN_MAC_NAME:=$(APP)-$(GOOS_MAC)-$(GOARCH_MAC)
BIN_LINUX_NAME:=$(APP)-$(GOOS_LINUX)-$(GOARCH_LINUX)

ANALYZE_THRESHOLD:=10
ANALYZE_MAX_LENGTH:=80

default: analyze test benchmark

format:
	@go fmt $$(go list ./... | grep -v /vendor/)

analyze:
	@echo "The following tools are required: vet, goconst, gosec, usedexports, lll, prealloc, ineffassign, misspell, and deadcode"; echo;

	@# Examines Go source code and reports suspicious constructs, such as Printf calls
	@#   whose arguments do not align with the format string.
	@echo "Running vet"
	@go vet -all $$(go list ./... | grep -v /vendor/)

	@# Find repeated strings that could be replaced by a constant.
	@echo "Running goconst"
	@goconst -ignore "vendor" $$(go list ./... | grep -v /vendor/)

	@# Inspect source code for security problems by scanning the Go AST.
	@echo "Running gosec"
	@gosec ./...

	@# Find exported variables (const, var, func, struct) in Go that could be unexported.
	@echo "Running usedexports"
	@usedexports .

	@# Enforce line length in files (120)
	@echo "Running lll"
	@lll --maxlength $(ANALYZE_MAX_LENGTH) *.go

	@# Find slice declarations that could potentially be preallocated.
	@echo "Running prealloc"
	@prealloc -set_exit_status .

	@# Detect ineffectual assignments in Go code.
	@echo "Running ineffassign"
	@ineffassign .

	@# Check for deadcode
	@find . -type d -path "*" | xargs deadcode

	@# Check english
	@echo "Running misspell"
	@misspell -w *.go

benchmark:
	@go test -run=XYZ -benchmem -bench .

test: format
	@go test -v -cover -race

docs:
	@echo "Navigate to http://localhost:6060/pkg/git.cto.ai/thalesfsp/ops/pkg/"
	@godoc -http=":6060"

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

.PHONY: format
	analyze
	benchmark
	test
	docs
	list
