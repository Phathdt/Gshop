FROM golang:1.17.5-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o gshop ./cmd/migrate/main.go

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=builder --chown=nobody:nobody ./app/gshop /app/
COPY --from=builder --chown=nobody:nobody ./app/migrations /app/migrations/
COPY --from=builder --chown=nobody:nobody ./app/cmd/migrate/run.sh .

ENTRYPOINT sh run.sh
