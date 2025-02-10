FROM golang:1.19-alpine

COPY main.go /app/main.go

WORKDIR /app

RUN go build -o main main.go
CMD ["sh", "-c", "./main > output.txt"]
