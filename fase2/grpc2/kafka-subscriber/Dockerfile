FROM golang:1.16-alpine

RUN apk update && apk add bash

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /subscriber

EXPOSE 3000

CMD [ "/subscriber" ]