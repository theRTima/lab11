# ─── Stage 1: Build ───────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o app \
    ./cmd/main.go

FROM scratch

COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]