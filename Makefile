BIN_NAME = t
INSTALL_DIR = $(HOME)/.local/bin
FISH_COMP_DIR = $(HOME)/.config/fish/completions

.PHONY: build test install uninstall clean

build:
	go build -o $(BIN_NAME) .

test:
	go test ./...

install: build
	mkdir -p $(INSTALL_DIR) $(FISH_COMP_DIR)
	cp $(BIN_NAME) $(INSTALL_DIR)/$(BIN_NAME)
	$(INSTALL_DIR)/$(BIN_NAME) completion fish > $(FISH_COMP_DIR)/$(BIN_NAME).fish

uninstall:
	rm -f $(INSTALL_DIR)/$(BIN_NAME)
	rm -f $(FISH_COMP_DIR)/$(BIN_NAME).fish

clean:
	rm -f $(BIN_NAME)
