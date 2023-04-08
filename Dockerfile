FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY .env firebase_credential.json server.crt server.key email.html ./
COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
