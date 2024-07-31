FROM --platform=linux/amd64 golang:1.22.2-alpine AS builder

RUN apk add bash ca-certificates gcc musl-dev

WORKDIR /build
COPY . .
RUN go mod download

RUN go build -tags musl -ldflags '-extldflags "-static"' -o /build/main ./cmd/main/main.go

FROM --platform=linux/amd64 scratch

WORKDIR /app
COPY --from=builder /build/main .

EXPOSE 8080
EXPOSE 44049

ENTRYPOINT ["/app/main"]