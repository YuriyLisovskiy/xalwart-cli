MKDIR_P = mkdir -p
RM = rm -rf

APP_NAME = xalwart

.PHONY: clean build compile-single compile run install-unix install-win

all: clean build compile run

clean:
	@${RM} bin/*

build:
	@echo "Compiling target..."
	@${MKDIR_P} bin
	@${RM} bin/${APP_NAME}
	@go build -o bin/${APP_NAME} src/cmd/main.go
	@echo "Done."

run:
	@go run src/cmd/main.go

compile-single:
	@echo "Compiling for $(OS)($(ARCH))..."
	@${RM} rm -rf -- bin/$(OS)-$(ARCH)
	@${MKDIR_P} bin/$(OS)-$(ARCH)
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$(OS)-$(ARCH)/${APP_NAME}${EXT} src/cmd/main.go
	@echo "Done.\n"

compile:
	@make -s compile-single OS=freebsd ARCH=386
	@make -s compile-single OS=linux ARCH=386
	@make -s compile-single OS=linux ARCH=arm
	@make -s compile-single OS=linux ARCH=arm64
	@make -s compile-single OS=windows ARCH=386 EXT=.exe
	@make -s compile-single OS=windows ARCH=amd64 EXT=.exe

install-unix:
	cp bin/${APP_NAME} /usr/local/bin
	chmod a+x /usr/local/bin/${APP_NAME}

install-win:
	mkdir -p C:\${APP_NAME}
	cp bin\${APP_NAME}.exe C:\${APP_NAME}
	pathman /au C:\${APP_NAME}
