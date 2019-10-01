FROM golang:latest

LABEL maintainer="Pablo Perez Mollo <pablo.perez-mollo@hotmail.com>"

COPY . /go/src/github.com/pablitopm/go-minesweeper

WORKDIR /go/src/github.com/pablitopm/go-minesweeper

RUN go get -u ./...
RUN go build -o minesweeper cmd/minesweeper/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./minesweeper"]