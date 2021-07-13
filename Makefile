GO = go
GOINSTALL = $(GO) install
GOTEST = $(GO) test
GOBENCH = $(GOTEST) -bench=. -run=XXX -benchmem

LDFLAGS = -ldflags="-s -w"

CMDS = ./cmd/...
PKGS = ./pkg/...

install:
	$(GOINSTALL) $(LDFLAGS) $(CMDS)

test:
	$(GOTEST) $(PKGS)

bench:
	$(GOBENCH) $(PKGS)
