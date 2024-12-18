### THANK U KUBE-VIP

ifneq (,$(wildcard ./.env))
    include .env
endif

SHELL := /bin/sh

# The name of the executable
TARGET := ano

CSV ?= malstat.csv
DB ?= ${DATABASE}
REPOSITORY ?= reidaa
DOCKERFILE ?= build/Dockerfile
DOCKERTAG ?= latest

.PHONY: all build clean install uninstall check run deploy ansible

all: build

$(TARGET):
	go build -o $@ main.go

build: $(TARGET)
	@true

clean:
	rm -f $(TARGET)

re: clean build
.PHONY: re

run: build
	./$(TARGET) scrap --top 100 --db $(DB)

run-help: build
	./$(TARGET) help

run-serve: build
	./$(TARGET) serve

ansible:
	ansible-playbook deployments/ansible/deploy.yml -vv 

deploy: build ansible clean

lint: build
	golangci-lint run

docker:
	docker build -t ${REPOSITORY}/${TARGET}:${DOCKERTAG} -f ${DOCKERFILE} .
.PHONY: docker

docker-build-debug:
	docker build --progress=plain --no-cache -t ${REPOSITORY}/${TARGET}:debug -f ${DOCKERFILE} .
.PHONY: docker-build-debug

docker-run: docker
	docker run -p 8080:8080 ${REPOSITORY}/${TARGET}:${DOCKERTAG} version
.PHONY: docker-run