FROM golang:1.24 AS builder

COPY . /app
WORKDIR /app

RUN ["go", "mod", "download"]
RUN ["go", "build", "-o", "api", "."]

FROM alpine:3.21

COPY --from=builder /app/api /app/api

ENTRYPOINT ["/app/api"]
