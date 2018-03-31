VERSION:=$(shell git describe --tags)
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

ifeq ($(PREFIX),)
	PREFIX := /usr/local
endif

plugin-png.so: src
	go build -buildmode=plugin $(LDFLAGS) -o plugin-png.so ./src

.PHONY: deps
deps:
	dep ensure

install:
	install -d $(DESTDIR)$(PREFIX)/lib/vexlink/plugins/
	install -m 644 plugin-png.so $(DESTDIR)$(PREFIX)/lib/vexlink/plugins/

clean:
	go clean

