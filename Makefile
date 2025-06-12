.PHONY: build
build:
	docker build --platform linux/amd64 -t ghcr.io/unhanded/txp:dev .

.PHONY: dev
dev:
	PORT=8089 TXP_DIR=./docs/example go run ./app/txpd

dist/amd64:
	GOARCH=amd64 go build -o ./dist/amd64/txpd ./app/txpd

dist/arm64:
	GOOS=darwin;GOARCH=arm64 go build -o ./dist/arm64/txpd ./app/txpd
