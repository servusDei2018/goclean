GO=go build -ldflags="-s -w"

all:
	$(GO)

install:
	go install
