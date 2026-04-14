FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN go build -o /go/bin/app ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app"]

LABEL Name=moviemanagement Version=0.0.1
EXPOSE 8080