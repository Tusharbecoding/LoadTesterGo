FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o load-tester-go ./cmd/load-tester-go

CMD ["./load-tester-go"]
