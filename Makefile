### THANK U KUBE-VIP

SHELL := /bin/sh

TARGET := scrapper
CSV := malstat.csv

# These will be provided to the target
BUILD := `git rev-parse HEAD`


# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Build=$(BUILD)"

.PHONY: all build clean install uninstall check run

all: check install

$(TARGET):
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	rm -f $(TARGET)
	rm -f $(CSV)

install:
	@go install $(LDFLAGS)

uninstall: clean
	rm -f $$(which ${TARGET})

check:
	go mod tidy

run: install
	@$(TARGET) scrap --top 100 --csv $(CSV)