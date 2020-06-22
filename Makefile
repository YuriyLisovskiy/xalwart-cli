MKDIR_P = mkdir -p
RM = rm -rf

APP_NAME = xalwart

.PHONY: clean build compile-single compile run

all: clean build compile run

clean:
	@${RM} bin/*

build:
	@echo "Compiling target..."
	@${RM} bin/${APP_NAME}
	@go build -o bin/${APP_NAME} cmd/main.go
	@echo "Done."

run:
	@go run cmd/main.go

compile-single:
	@echo "Compiling for $(OS)($(ARCH))..."
	@${RM} rm -rf -- bin/$(OS)-$(ARCH)
	@${MKDIR_P} bin/$(OS)-$(ARCH)
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$(OS)-$(ARCH)/${APP_NAME} cmd/main.go
	@echo "Done.\n"

compile:
	@make -s compile-single OS=freebsd ARCH=386
	@make -s compile-single OS=linux ARCH=386
	@make -s compile-single OS=linux ARCH=arm
	@make -s compile-single OS=linux ARCH=arm64
	@make -s compile-single OS=windows ARCH=386
