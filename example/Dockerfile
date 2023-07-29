FROM golang:latest

COPY . /app
WORKDIR /app/example

RUN go build -o ../bin/client

CMD ["/app/bin/client"]
