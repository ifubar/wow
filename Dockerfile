FROM golang:latest

COPY . /app
WORKDIR /app
RUN go build -o bin/wow

CMD ["/app/bin/wow"]
