.PHONY: build
build:
	docker build -t unhanded/txp:latest .

.PHONY: dev
dev:
	PORT=8089 TXP_DIR=./example go run ./app/txpd
