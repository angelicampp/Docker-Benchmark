FROM golang:1.21-alpine AS builder
WORKDIR /

COPY main.go /main.go
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /main .
CMD ["sh", "-c", "./main > /app/output.txt"]
