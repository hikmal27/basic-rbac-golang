FROM golang:1.17-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go run main.go db:migrate
RUN go build -o main .

CMD ["/app/main"]