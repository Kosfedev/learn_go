# Стадия сборки
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o learn_go_be ./cmd/main.go

FROM alpine:latest

# TODO: протестировать без
RUN apk --no-cache add tzdata

RUN adduser -D appuser
USER appuser

WORKDIR /app

COPY --from=builder --chown=appuser:appuser /app/learn_go_be ./
COPY --from=builder --chown=appuser:appuser /app/.env.prod ./
COPY --from=builder --chown=appuser:appuser /app/pkg/docs ./pkg/docs/

# TODO: протестировать без
ENV GRPC_PORT=50051
ENV GRPC_GW_PORT=8080
ENV MODE="prod"

EXPOSE $GRPC_PORT $GRPC_GW_PORT

CMD ["./learn_go_be"]