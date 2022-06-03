FROM golang:1.18.2-alpine3.15

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main .

CMD ["./main"]