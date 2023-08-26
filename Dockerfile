# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder
WORKDIR /github.com/ddddddO/gtree
COPY go.* *.go ./
WORKDIR /github.com/ddddddO/gtree/markdown
COPY markdown/*.go ./
WORKDIR /github.com/ddddddO/gtree/cmd/gtree
COPY cmd/gtree/*.go ./
RUN go build -o gtree .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /github.com/ddddddO/gtree/cmd/gtree/gtree ./
ENTRYPOINT ["./gtree"]
