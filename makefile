# ---- Go-specific helpers ----------------------------------------------------
GOFILES     := $(shell find . -name '*.go' -type f)
GOMOD       := go.mod
GOSUM       := go.sum

DIST        := lxtxt
BINARY      := dist/$(DIST)

.PHONY: install upgrade uninstall clean coverage

# ---- Installation helpers (change if necessary) -----------------------------
INSTALL_DIR := /usr/local/bin/

# ---- Default target ---------------------------------------------------------
all: $(BINARY)

# ---- Build rule -------------------------------------------------------------
$(BINARY): $(GOFILES) $(GOMOD) $(GOSUM)
	mkdir -p dist/
	go build -o $@ ./cmd/lxtxt

# ---- Install / upgrade / uninstall / clean ----------------------------------
install: upgrade

upgrade: $(BINARY)
	sudo cp $(BINARY) /usr/local/bin/

uninstall:
	sudo rm -rf /usr/local/bin/$(DIST)

clean:
	rm -rf dist

coverage:
	@go test ./... -coverprofile=.tmp.out
	@go tool cover -html=.tmp.out
	@rm .tmp.out
