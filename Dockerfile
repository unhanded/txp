ARG platform='amd64'

FROM golang:1.24.3-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o ./dist/txpd ./app/txpd

FROM --platform=${platform} ghcr.io/unhanded/typisch:latest

COPY --from=builder /app/dist/txpd /usr/bin/txpd

ENV TXP_DIR=/txp_data

VOLUME [ "/txp_data" ]
VOLUME [ "/fonts" ]

ENTRYPOINT [ "/bin/sh" ]
