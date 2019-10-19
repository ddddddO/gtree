FROM golang:1.13.2-alpine

WORKDIR /sub
COPY /app/sub /sub
COPY /Dockerfiles/app/wait-for-it.sh /sub

ENV GO111MODULE on

ENTRYPOINT ["./wait-for-it.sh"]
CMD ["go", "run", "main.go"]
