FROM golang:1.12.0-alpine3.9

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY entrypoint.sh /entrypoint.sh
RUN go build -o /annotator
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]