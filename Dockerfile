# Étape de construction
FROM golang:1.22-alpine

WORKDIR /test_http

LABEL maintainer="Arocchet <https://zone01normandie.org/git/arocchet/>"
LABEL version="1.0.0"
LABEL description="This is a Go application that serves HTTP requests."

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main && \
    apk update && \
    apk add bash

EXPOSE 8080

CMD ["./main"]