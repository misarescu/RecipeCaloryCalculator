BUILD_DIR = bin

build:
	@go build -o ./$(BUILD_DIR)/ ./cmd/...

clean:
	@rm -rf $(BUILD_DIR)