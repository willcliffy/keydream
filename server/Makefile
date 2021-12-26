.PHONY: lobby gameserver mockclient

up-build:
	docker-compose up --build

up:
	docker-compose up


lobby-build:
	go build -o ./_dist/lobby /app/cmd/lobby/main.go

lobby:
	go run cmd/lobby/main.go


gameserver-build:
	go build -o ./_dist/gameserver /app/cmd/gameserver/main.go

gameserver:
	go run cmd/gameserver/main.go


mockclient:
	go run cmd/mock_client/main.go
