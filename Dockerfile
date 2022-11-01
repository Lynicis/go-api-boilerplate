FROM golang:1.19.2-alpine

WORKDIR /usr/src/app

RUN chmod +x go-api-boilerplate

EXPOSE 8080

CMD ./go-rest-api-boilerplate