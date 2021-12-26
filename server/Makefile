.PHONY: lobby world mockclient

up-build:
	docker-compose up --build

up:
	docker-compose up


lobby-build:
	go build -o ./_dist/lobby /app/cmd/lobby/main.go

lobby:
	go run cmd/lobby/main.go


world-build:
	go build -o ./_dist/world /app/cmd/world/main.go

world:
	go run cmd/world/main.go


mockclient:
	go run cmd/mock_client/main.go
