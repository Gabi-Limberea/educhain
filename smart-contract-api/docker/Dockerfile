FROM golang:1.24-alpine3.21 AS builder

COPY ./smart-contract-api /app
COPY ./ipfs-api /ipfs-api
WORKDIR /app

RUN ["go", "mod", "download"]
RUN ["go", "build", "-o", "api", "."]

FROM alpine:3.21

COPY --from=builder /app/api .

ENTRYPOINT ["./api"]
