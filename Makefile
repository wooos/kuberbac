BINDIR      := $(CURDIR)/bin
BINNAME 	?= kuberbac

.PHONY build:
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME):
	go build -o $(BINDIR)/$(BINNAME) main.go

.PHONY clean:
clean:
	@rm -rf bin/
