FROM golang:alpine as builder

RUN apk update && apk add --no-cache ca-certificates git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ADD app/setup.go.prod app/setup.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /src/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY .env firebase_credential.json sheets_credential.json email.html ./

VOLUME ["/cert-cache"]

ENTRYPOINT ["./main"]
