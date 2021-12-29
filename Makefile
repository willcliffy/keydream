.PHONY: lobby world mockclient client

server-up-build:
	cd server && make up-build

server-up:
	cd server && make up


server-lobby-build:
	cd server && make lobby-build

server-lobby:
	cd server && make lobby


server-world-build:
	cd server && make world-build

server-world:
	cd server && make world


mockclient:
	cd server && make mockclient

client:
	love ./client
