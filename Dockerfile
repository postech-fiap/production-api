FROM golang:1.21.6-alpine3.19

WORKDIR /app

COPY . .

RUN rm main
RUN go build cmd/http/main.go

CMD ["./main"]

EXPOSE 8080
