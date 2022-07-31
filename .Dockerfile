FROM golang:1.18

COPY . ../tic-tac-toe

RUN go version
RUN go build . 
RUN go run ./

ENTRYPOINT "./_bin/tic-tac-toe"
