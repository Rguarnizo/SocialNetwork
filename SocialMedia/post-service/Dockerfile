FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /app/bin
RUN go build -o /app/bin/post-service ./post-service/cmd

EXPOSE 8080

CMD ["/app/bin/post-service"]