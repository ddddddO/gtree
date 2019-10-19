FROM golang:1.13.2-alpine

WORKDIR /pubsub
COPY /app/pubsub /pubsub
COPY /Dockerfiles/app/wait-for-it.sh /pubsub

ENV GO111MODULE on

ENTRYPOINT ["./wait-for-it.sh"]
CMD ["go", "run", "main.go"]
