VERSION:=$(git describe --tags)
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

ifeq ($(PREFIX),)
	PREFIX := /usr/local
endif

plugin-png.so:
	go build -buildmode=plugin $(LDFLAGS) -o plugin-png.so

install: plugin-png.so
	install -d $(DESTDIR)$(PREFIX)/lib/vexlink/plugins/
	install -m 644 plugin-png.so $(DESTDIR)$(PREFIX)/lib/vexlink/plugins/

clean:
	go clean

