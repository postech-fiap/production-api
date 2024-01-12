FROM golang:1.21.6-alpine3.19

WORKDIR /app/build

COPY . .

RUN go get ./...
RUN go build cmd/http/main.go
RUN mv main resources ../

WORKDIR /app

RUN rm -rf build

CMD ["./main"]

EXPOSE 8080
