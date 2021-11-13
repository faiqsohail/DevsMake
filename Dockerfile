# Install dependencies only when needed
FROM node:alpine AS deps
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

# Rebuild the source code only when needed
FROM node:alpine AS builder
WORKDIR /app
COPY frontend/. .
COPY --from=deps /app/node_modules ./node_modules
ENV NODE_OPTIONS=--openssl-legacy-provider
RUN yarn build && yarn install --production --ignore-scripts --prefer-offline

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

# Production image, copy all the files and run next
FROM node:alpine AS runner
WORKDIR /app

RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001

COPY --from=go-builder /go/src/devsmake/main ./main
COPY --from=go-builder /go/src/devsmake/startup.sh ./startup.sh
COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

USER nextjs

EXPOSE 80 8080

CMD ["sh", "startup.sh"]