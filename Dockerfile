FROM golang:1.17.5-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o app ./main.go

# ----------------------
FROM ghcr.io/redocly/redoc/cli:v2.0.0-rc.59 as redoc
WORKDIR /app
COPY docs.yaml ./
RUN redoc-cli bundle docs.yaml

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=builder --chown=nobody:nobody ./app /app/
COPY --from=builder --chown=nobody:nobody /app/run.sh .
COPY --from=redoc --chown=nobody:nobody /app/redoc-static.html ./public/index.html

ENTRYPOINT sh run.sh
