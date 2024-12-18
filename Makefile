### THANK U KUBE-VIP

ifneq (,$(wildcard ./.env))
    include .env
endif

SHELL := /bin/sh

# The name of the executable
TARGET := malstatback

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
