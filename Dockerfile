FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go get ./...

COPY . .

RUN go build -o go-rest-api-boilerplate cmd/main.go

EXPOSE 8080

CMD ./go-rest-api-boilerplate