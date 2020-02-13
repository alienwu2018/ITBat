FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/ITBat
COPY . $GOPATH/src/github.com/ITBat
RUN go build -o server
EXPOSE 8000
ENTRYPOINT ["./server"]