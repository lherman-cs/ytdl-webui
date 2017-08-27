FROM arm32v6/golang:alpine

WORKDIR /go/src/app

RUN apk update &&\
    apk add --no-cache git &&\
    git clone https://github.com/lherman-cs/ytdl-webui.git . &&\
    go-wrapper download &&\
    go-wrapper install &&\
    apk del git
CMD ["go-wrapper", "run"]
