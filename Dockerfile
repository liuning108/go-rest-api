FROM golang:1.22 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest as  production
COPY --from=builder /app .
CMD ["./app"]