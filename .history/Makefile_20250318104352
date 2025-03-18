.PHONY: left-right cal-path server all

left-right:
	go run cmd/left_right/main.go

cal-path:
	go run cmd/cal_path/main.go

server:
	go run cmd/main_clause/main.go

all: left-right cal-path server

APP_NAME = max-path
SRC = main.go
run:
	go run $(SRC)

build:
	go build -o $(APP_NAME)

start:
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)
test:
	go test -v ./...