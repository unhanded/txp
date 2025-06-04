FROM golang:1.24.3-bookworm as builder

WORKDIR /app

COPY . .

RUN go build -o ./dist/txpd ./app/txpd

FROM ghcr.io/unhanded/typisch:v0.13.1a

COPY --from=builder /app/dist/txpd /usr/bin/txpd

ENV TXP_DIR=/txp_data

VOLUME [ "/txp_data" ]

ENTRYPOINT [ "txpd" ]