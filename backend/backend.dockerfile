FROM golang:latest

WORKDIR /app/backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env /app/backend

WORKDIR /app/backend/cmd/server

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
