BINDIR      := $(CURDIR)/bin
BINNAME 	?= kuberbac

GOPATH        = $(shell go env GOPATH)

# go option
PKG        := ./...
TAGS       :=
TESTS      := .
TESTFLAGS  :=
LDFLAGS    := -w -s
GOFLAGS    :=
SRC        := $(shell find . -type f -name '*.go' -print)

# ---------------------------------------------------------------------------------------------------------------------
# build
.PHONY build:
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME):
	GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME) ./cmd/kuberbac/

#
.PHONY clean:
clean:
	@rm -rf bin/
