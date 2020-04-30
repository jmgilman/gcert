FROM golang:1.14-alpine AS staging

RUN mkdir /build
COPY . /build/
COPY proto /build/proto

WORKDIR /build
RUN go build .
RUN chmod +x gcert

FROM alpine:latest

RUN addgroup -S appuser && adduser -S appuser -G appuser
USER appuser
WORKDIR /opt

COPY --from=staging --chown=appuser /build/gcert .
CMD ["/opt/gcert"]