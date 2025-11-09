BINARY_NAME = gocp

INSTALL_DIR = $(HOME)/.local/bin

build:
	go build -o $(BINARY_NAME) ./cmd/go-colors-picker


install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)"


clean:
	rm -f $(BINARY_NAME)
