FROM golang:latest

LABEL maintainer="Pablo Perez Mollo <pablo.perez-mollo@hotmail.com>"

COPY . /go/src/github.com/pablitopm/go-minesweeper

WORKDIR /go/src/github.com/pablitopm/go-minesweeper

RUN go get ./...
RUN go build -o minesweeper cmd/minesweeper/main.go

EXPOSE 8080

CMD ["./minesweeper"]

