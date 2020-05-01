FROM golang:1.14-alpine AS staging

ARG VAULTBOT_REPO=https://gitlab.com/msvechla/vaultbot
ARG VAULTBOT_VER=v1.9.0

RUN apk update && apk upgrade && apk add git vault

# Build VaultBot
RUN mkdir -p /build/vaultbot
WORKDIR /build/vaultbot
RUN git clone -b $VAULTBOT_VER $VAULTBOT_REPO .
RUN go build .

# Build gcert service
RUN mkdir /build/gcert
COPY . /build/gcert
COPY files /build/gcert/files
COPY proto /build/gcert/proto

WORKDIR /build/gcert
RUN go build .


FROM alpine:latest

RUN apk update && apk upgrade && apk add bash
RUN addgroup -S appuser && adduser -S appuser -G appuser
USER appuser
WORKDIR /opt

COPY --from=staging --chown=appuser /usr/sbin/vault .
COPY --from=staging --chown=appuser /build/vaultbot/vaultbot .
COPY --from=staging --chown=appuser /build/gcert/gcert .
COPY --from=staging --chown=appuser /build/gcert/files/init.sh .
RUN chmod +x init.sh
ENTRYPOINT ["/opt/init.sh"]