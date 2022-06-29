FROM golang:1.18-alpine

WORKDIR /usr/src/app

RUN chmod +x go-api-boilerplate

EXPOSE 8080 8081

CMD ./go-rest-api-boilerplate