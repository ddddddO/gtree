FROM golang:1.15-alpine AS builder

WORKDIR /cmd
COPY go.* ./
COPY ./cmd/api/main.go .
RUN go build -o api main.go

FROM alpine
COPY --from=builder /cmd/api .

CMD ["./api"]