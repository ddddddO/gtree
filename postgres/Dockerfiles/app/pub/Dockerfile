FROM golang:1.13.2-alpine

WORKDIR /pub
COPY /app/pub /pub
COPY /Dockerfiles/app/wait-for-it.sh /pub

ENV GO111MODULE on

ENTRYPOINT ["./wait-for-it.sh"]
CMD ["go", "run", "main.go"]
