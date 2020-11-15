FROM golang:1.13.0-alpine3.10

WORKDIR go/src/github.com/hoomnayangi/api

COPY . .

RUN  mkdir bin

RUN go build -o bin/server cmd/*.go

RUN rm -f .env

RUN mv .env.dev .env

CMD bin/server