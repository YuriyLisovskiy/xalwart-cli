RM = rm -rf
APP_NAME = xalwart

.PHONY: clean build run install

all: clean build

clean:
	@${RM} bin/*

build:
	@echo "Compiling target..."
	@mkdir -p bin
	@${RM} bin/${APP_NAME}
	@go build -o bin/${APP_NAME} cli/main.go
	@echo "Done."

run:
	@go run cli/main.go

install:
	cp bin/${APP_NAME} /usr/local/bin
	chmod a+x /usr/local/bin/${APP_NAME}
