FROM golang:1.22.2-alpine AS builder

RUN apk add --no-progress --no-cache gcc musl-dev

WORKDIR /build
COPY . .
RUN go mod download

RUN go build -tags musl -ldflags '-extldflags "-static"' -o /build/main ./cmd/main/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /build/main .

EXPOSE 8080
EXPOSE 44049

ENTRYPOINT ["/app/main"]