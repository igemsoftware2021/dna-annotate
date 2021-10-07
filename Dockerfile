FROM golang:1.12.0-alpine3.9

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /main
RUN chmod +x /entrypoint.sh
CMD ["/entrypoint.sh"]