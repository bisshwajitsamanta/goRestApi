FROM golang:1.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=Linux go build -o app cmd/server/main.go

FROM alipine:latest AS production
COPY --from=builder /app .
CMD ["./app"]