MKDIR_P = mkdir -p
RM = rm -rf

.PHONY: clean build compile-single compile run

all: clean build compile run

clean:
	@${RM} bin/*

build:
	@echo "Compiling target..."
	@${RM} bin/wasp
	@go build -o bin/wasp cmd/main.go
	@echo "Done."

run:
	@go run cmd/main.go

compile-single:
	@echo "Compiling for $(OS)($(ARCH))..."
	@${RM} rm -rf -- bin/$(OS)-$(ARCH)
	@${MKDIR_P} bin/$(OS)-$(ARCH)
	@GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/$(OS)-$(ARCH)/wasp cmd/main.go
	@echo "Done.\n"

compile:
	@make -s compile-single OS=freebsd ARCH=386
	@make -s compile-single OS=linux ARCH=386
	@make -s compile-single OS=linux ARCH=arm
	@make -s compile-single OS=linux ARCH=arm64
	@make -s compile-single OS=windows ARCH=386
