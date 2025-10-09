FROM golang:1.25.1-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -C /app/cmd/sso -o sso


FROM alpine:3.22.2
WORKDIR /root/
COPY --from=builder /app ./app
EXPOSE 3000
WORKDIR /root/app
CMD ["./cmd/sso/sso"]
