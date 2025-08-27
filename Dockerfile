ARG platform='amd64'

FROM --platform=${platform} golang:1.24.3-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o ./dist/txpd ./app/txpd

FROM --platform=${platform} ghcr.io/unhanded/typisch:v0.13.1.7

COPY --from=builder /app/dist/txpd /usr/bin/txpd

ENV TXP_DIR=/txp_data

VOLUME [ "/txp_data" ]
VOLUME [ "/fonts" ]

ENTRYPOINT [ "/bin/sh" ]
