FROM golang:1.23
WORKDIR /src
RUN go build -o /bin/hello ./main.go
