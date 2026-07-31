FROM golang:1.26.5
WORKDIR /app
COPY go.mod go.sum .
RUN go get ./...
COPY */*.go .
RUN go build -o ./DemoAuthService ./...
CMD ["./DemoAuthService"]