APPLICATION := $(lastword $(subst /, ,$(dir $(CURDIR))))
PLATFORM := $(shell go env GOOS)_$(shell go env GOARCH)
PACKAGE := $(shell go list)
TESTS := $(wildcard *_test.go **/*_test.go)
SRC := $(filter-out $(TESTS), $(wildcard *.go **/*.go))
BIN := $(value GOPATH)\pkg\$(PLATFORM)\$(PACKAGE).a

echo: 
	@echo $(value BIN)

# build when changed
$(BIN): $(SRC)
	go install

# run command at fixed intervals
watch: install
	@kouhai -i 2s 'make test'

# run all tests and rebuild when changed
test: $(BIN)
	@go test $(PACKAGE)/...

# build and install application
install: $(BIN)

# remove application
clean: 
	@go clean -i

.PHONY: clean test watch install
