FROM golang:1.23-bullseye AS builder
WORKDIR /src
ENV GOCACHE=/root/.cache/go-build \
    CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=bind,target=. \
    go build -o /app/binary

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/binary .
COPY config*.toml .
EXPOSE 8080
CMD ["./binary"]