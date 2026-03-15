.PHONY: build install build-all build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-windows-arm64 clean

BINARY_NAME := gofm
BUILD_DIR := build
BIN_DIR ?= $(HOME)/.local/bin

build:
	@mkdir -p $(BUILD_DIR) && go build -o $(BUILD_DIR)/$(BINARY_NAME) .

install: build
	@mkdir -p $(BIN_DIR)
	@install -m 755 $(BUILD_DIR)/$(BINARY_NAME) $(BIN_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(BIN_DIR)/$(BINARY_NAME)"

build-linux-amd64:
	@mkdir -p $(BUILD_DIR) && GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
build-linux-arm64:
	@mkdir -p $(BUILD_DIR) && GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .

build-darwin-amd64:
	@mkdir -p $(BUILD_DIR) && GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
build-darwin-arm64:
	@mkdir -p $(BUILD_DIR) && GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

build-windows-amd64:
	@mkdir -p $(BUILD_DIR) && GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
build-windows-arm64:
	@mkdir -p $(BUILD_DIR) && GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe .



build-all: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-windows-arm64
	@mkdir -p $(BUILD_DIR)
	@echo "Built all platforms in $(BUILD_DIR)/"

clean:
	@rm -rf $(BUILD_DIR)
	@echo "Removed $(BUILD_DIR)/"