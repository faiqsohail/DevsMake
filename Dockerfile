# Install dependencies only when needed
FROM node:alpine AS js-deps
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

# Rebuild the source code only when needed
FROM node:alpine AS js-builder
WORKDIR /app
COPY . .
COPY --from=js-deps /app/node_modules ./node_modules
RUN npm run build

# Build go image
FROM golang:1.17.1-alpine AS go-builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/devsmake
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build cmd/devs-make-server/main.go

# Bring it all together on an alpine image
FROM alpine:3.13.6 AS runner
WORKDIR /webapp
COPY --from=js-builder /app/dist ./dist
COPY --from=go-builder /go/src/devsmake/tls ./tls
COPY --from=go-builder /go/src/devsmake/main ./main

EXPOSE 5000 5001

CMD ["./main","--tls-certificate=tls/tls.cert","--tls-key=tls/tls.key","--host=0.0.0.0","--tls-port=5000", "--port=5001"]