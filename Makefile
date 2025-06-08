.PHONY: build
build:
	docker build --platform linux/amd64 -t ghcr.io/unhanded/txp:dev .

.PHONY: dev
dev:
	PORT=8089 TXP_DIR=./example go run ./app/txpd
