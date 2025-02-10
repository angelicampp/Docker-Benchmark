FROM golang:1.19-alpine

COPY main.go /main.go

RUN go build -o main main.go
CMD ["sh", "-c", "./main > /app/output.txt"]
