FROM golang:1.18.4-alpine3.16

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out ./
ENTRYPOINT ["/out"]