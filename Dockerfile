FROM golang:1.24.5-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
