FROM golang:1.16 AS builder

WORKDIR /app

COPY ./ .

RUN CGO_ENABLED=0 go build -o server-2 main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server-2 ./

RUN apk --no-cache add tzdata

CMD [ "./server-2" ]
