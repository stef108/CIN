FROM golang:1.26rc2-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# CGO_ENABLED=0 creates a statically linked binary (no external C library dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./main"]
