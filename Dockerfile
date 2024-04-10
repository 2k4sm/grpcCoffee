FROM golang:1.21-alpine3.18 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o httpcoffee .

FROM golang:1.21-alpine3.18 as runner
WORKDIR /app
COPY --from=builder /app/httpcoffee .
EXPOSE 6969

CMD ["./httpcoffee"]