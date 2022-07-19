# FROM golang:alpine
# 
# WORKDIR /app
# 
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
# 
# COPY . .
# 
# RUN go build -o /gRPC-client-api
# 
# EXPOSE 8080
# 
# CMD [ "/gRPC-client-api" ]

FROM golang:latest AS apigo

RUN apt-get update

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src

COPY go.mod .
RUN go mod download
COPY . .
RUN go build gRPC-client.go

FROM scratch
COPY --from=apigo /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=apigo /go/src .

EXPOSE 8080

ENTRYPOINT [ "./gRPC-client" ]