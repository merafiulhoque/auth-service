# build stage
FROM golang:1.26-alpine as builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/server ./cmd/server/main.go

# production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/server /server

EXPOSE 2110

CMD [ "/server" ]