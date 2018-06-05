# This how we want to name the binary output
BINARY=hotbed
# These are the values we want to pass for VERSION  and BUILD
VERSION=`git describe --tags`
TIME=`date +%FT%T%z`
MODE=$(mode)
ifeq ($(MODE),)
	MODE=dev
endif
# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.BuildVersion=${VERSION} -X main.BuildTime=${TIME} -X main.BuildMode=${MODE}"
# Builds the project

.PHONY: build
build:
	@echo "build..."
	go build ${LDFLAGS} -o ${BINARY}

.PHONY: clean
clean:
	@echo "clean..."
	rm -rf ${BINARY}
	rm -rf dist

.PHONY: deps
deps:
	@echo "deps..."
	dep ensure -update

.PHONY: dist
dist: clean build
	@echo "dist..."
	mkdir dist
	cp ${BINARY} dist/
	mkdir dist/configs
	cp configs/config_"${MODE}".ini dist/configs/
	cp docs dist/

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./.