.DEFAULT_GOAL := help

.git/hooks/pre-commit:
	@echo "make test" > $@
	@chmod +x $@

bin/transcript: main.go stt
	@go build -ldflags="-s -w" -o $@

bin/transcript64.exe: main.go stt
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $@

bin/transcript32.exe: main.go stt
	@GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o $@

releases/release.tar.gz: bin/transcript bin/transcript64.exe bin/transcript32.exe
	@mkdir -p releases/
	@tar --exclude '*.tar.gz' -zcf $@ bin

.PHONY: help
help:
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | sort | awk -F ':.*?## ' 'NF==2 {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: .git/hooks/pre-commit test ## setup development environment

.PHONY: test
test: ## test application
	@GOOGLE_APPLICATION_CREDENTIALS="$(PWD)/credentials.json" go test ./...

.PHONY: run
run: ## run application, use ARGS variable to send arguments. Usage: make run ARGS="gs://.../audio_sample.flac"
	@GOOGLE_APPLICATION_CREDENTIALS="$(PWD)/credentials.json" go run ./main.go $(ARGS)

.PHONY: build
build: releases/release.tar.gz ## Build bin folder and release package
