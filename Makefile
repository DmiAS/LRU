.PHONY: all
all: get run

.PHONY: get
get :
	go mod tidy

.PHONY: run
run:
	go run ./cmd/main.go
