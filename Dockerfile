# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

# Povlačenje zavisnosti
RUN go mod download

# POPRAVKA: Bildujemo samo glavni fajl, a ne "sve" (./...)
RUN go build -o /go/bin/app ./cmd/main.go

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Kopiramo binarni fajl iz builder-a
COPY --from=builder /go/bin/app /app

# Pokretanje
ENTRYPOINT ["/app"]

LABEL Name=moviemanagement Version=0.0.1
# Obrati pažnju: u handleru smo pominjali port 8080, 
# proveri da li ti se EXPOSE i port u kodu poklapaju.
EXPOSE 8080