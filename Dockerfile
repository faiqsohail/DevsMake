# Install dependencies only when needed and generate ssl cert
FROM node:alpine AS deps
WORKDIR /app
RUN apk add --no-cache libc6-compat && \
    apk add --no-cache openssl && \
    mkdir -p /app/tls && \
    openssl req -x509 -nodes -days 365 \
    -subj  "/C=CA/ST=AB/O=DevsMake/CN=devsmake.com" \
     -newkey rsa:2048 -keyout /app/tls/tls.key \
     -out /app/tls/tls.cert;
COPY package.json package-lock.json ./
RUN npm ci

# Rebuild the source code only when needed
FROM node:alpine AS js-builder
WORKDIR /app
COPY . .
COPY --from=deps /app/node_modules ./node_modules
RUN npm run build

# Build go image
FROM golang:1.17.1-alpine AS go-builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/devsmake
COPY . .
RUN go get -d -v ./... && \
    go install -v ./... && \
    go build cmd/devs-make-server/main.go

# Bring it all together on an alpine image
FROM alpine:3.13.6 AS runner
WORKDIR /webapp
COPY --from=deps /app/tls ./tls
COPY --from=js-builder /app/dist ./dist
COPY --from=go-builder /go/src/devsmake/main ./main
EXPOSE 443 80

CMD ["./main","--tls-certificate=tls/tls.cert","--tls-key=tls/tls.key","--host=0.0.0.0","--tls-port=443", "--port=80"]