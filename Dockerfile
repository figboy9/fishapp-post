FROM golang:1.13-alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine  
CMD ["./main"]
COPY --from=builder /app .