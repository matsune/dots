TARGET = dots
BUILD = ./build
OUT = $(BUILD)/$(TARGET)
CMD = ./cmd
GOBIN = ${GOPATH}/bin

.PHONY: build install clean test

build: 
	@go build -o $(OUT) $(CMD)

install: build
	@cp $(OUT) $(GOBIN)/$(TARGET)

clean:
	@-rm -rf $(BUILD)

test:
	@go test . -v -cover
	@go test $(CMD) -v -cover
