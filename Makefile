.PHONY: build

GO_FLAGS ?=
NAME     := mocktail
PACKAGE  := github.com/luisnquin/$(NAME)

install:
	@go install ./cmd/mocktail

build:
	@mkdir -p ./build
	@CGO_ENABLED=1 go build ${GO_FLAGS} \
		-ldflags "-w -s -a" \
		-o ./build/$(NAME) ./cmd/mocktail

race:
	@mkdir -p ./build
	@CGO_ENABLED=1 go build -race  -o ./build/$(NAME) ./cmd/mocktail
	@./build/$(NAME)

clean:
	@rm -rf ./build/*