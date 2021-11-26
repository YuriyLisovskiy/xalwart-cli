APP_NAME = xalwart
INSTALL_DIR = /usr/local/bin

.PHONY: clean build run install

all: clean build

clean:
	@rm -rf bin/*

build:
	@echo "Compiling $(APP_NAME)..."
	@mkdir -p bin
	@go build -o ./bin/$(APP_NAME) ./xalwart/main.go
	@echo "Done."

run:
	@go run ./xalwart/main.go

install:
	@echo "Installing $(APP_NAME) to $(INSTALL_DIR)"
	@cp bin/$(APP_NAME) $(INSTALL_DIR)
	@chmod a+x $(INSTALL_DIR)/$(APP_NAME)
	@echo "Done."
