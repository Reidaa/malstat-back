### THANK U KUBE-VIP

ifneq (,$(wildcard ./.env))
    include .env
endif

SHELL := /bin/sh

# The name of the executable (default is current directory name)
TARGET := scrapper

# These will be provided to the target
BUILD := `git rev-parse HEAD`

BIN_DIR ?= bin
CSV ?= malstat.csv
DB ?= ${DATABASE}
REPOSITORY ?= reidaa
DOCKERFILE ?= build/Dockerfile
DOCKERTAG ?= latest
# Use linker flags to provide version/build settings to the target
LDFLAGS := -ldflags "-X=main.Build=$(BUILD)"

init:
	go install mvdan.cc/gofumpt@latest

.PHONY: all build clean install uninstall check run deploy ansible

all: check install

$(TARGET):
	go build $(LDFLAGS) -o $(BIN_DIR)/$(TARGET)

build: $(TARGET)
	@true

clean:
	rm -rf $(BIN_DIR)

install:
	@go install $(LDFLAGS)

uninstall: clean
	rm -f $$(which ${TARGET})

run: install
	@$(TARGET) scrap --top 100 --csv $(BIN_DIR)/$(CSV) --db $(DB)

ansible:
	ansible-playbook deployments/ansible/deploy.yml -vv 

deploy: build ansible clean

lint: build
	golangci-lint run --enable-all --disable tagliatelle --disable wsl --disable varnamelen --disable exhaustruct --disable depguard 

format:
	gofumpt -l -w .

ci_check:
	go mod tidy
	test -z "$(git status --porcelain)"
	test -z $(shell gofmt -l *.go pkg cmd) || echo "[WARN] Fix formatting issues with 'make format'"
	golangci-lint run
	go vet ./...

docker-build:
	docker build -f $(DOCKERFILE) -t $(REPOSITORY)/$(TARGET):$(DOCKERTAG) .

docker-run: docker-build
	docker run --rm $(REPOSITORY)/$(TARGET):$(DOCKERTAG) scrap --top 100 --db $(DB)

docker-run-ghcr:
	docker run --rm ghcr.io/reidaa/malstat-scrapper:latest scrap --top 100 --db $(DB)