FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN apk update &&\
    apk add --no-cache git &&\
    go-wrapper download &&\
    go-wrapper install &&\
    apk del git

CMD ["go-wrapper", "run"]
