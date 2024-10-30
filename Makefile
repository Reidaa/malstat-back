### THANK U KUBE-VIP

ifneq (,$(wildcard ./.env))
    include .env
endif

SHELL := /bin/sh

# The name of the executable (default is current directory name)
TARGET := scrapper

# These will be provided to the target
BUILD := `git rev-parse HEAD`


REPOSITORY ?= reidaa
OUT_DIR = out
CSV := malstat.csv
DB := ${DATABASE}
# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Build=$(BUILD)"
DOCKERTAG ?= latest



.PHONY: all build clean install uninstall check run deploy

all: check install

$(TARGET):
	@go build $(LDFLAGS) -o $(OUT_DIR)/$(TARGET)

build: $(TARGET)
	@true

clean:
	rm -f $(OUT_DIR)/$(TARGET)
	rm -f $(OUT_DIR)/$(CSV)

install:
	@go install $(LDFLAGS)

uninstall: clean
	rm -f $$(which ${TARGET})

run: install
	@$(TARGET) scrap --top 100 --csv $(OUT_DIR)/$(CSV) --db $(DB)

deploy: build
	ansible-playbook deployments/ansible/deploy.yml -vv 
	
simplify:
	@gofmt -s -l -w *.go pkg cmd

check:
	go mod tidy
	test -z "$(git status --porcelain)"
	test -z $(shell gofmt -l *.go pkg cmd) || echo "[WARN] Fix formatting issues with 'make simplify'"
	golangci-lint run
	go vet ./...

docker-build:
	docker build -f build/Dockerfile -t $(REPOSITORY)/$(TARGET):$(DOCKERTAG) .

docker-run: docker-build
	docker run --rm $(REPOSITORY)/$(TARGET):$(DOCKERTAG) scrap --top 100 --db $(DB)