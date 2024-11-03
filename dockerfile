FROM golang:1.23.2-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /main .

CMD ["/main"]
