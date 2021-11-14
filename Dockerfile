# Install dependencies only when needed
FROM node:lts-alpine AS deps
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile

# Rebuild the source code only when needed
FROM node:lts-alpine AS builder
WORKDIR /app
COPY frontend/. .
COPY --from=deps /app/node_modules ./node_modules
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
FROM node:lts-alpine AS runner
WORKDIR /app

COPY --from=go-builder /go/src/devsmake/main ./main
COPY --from=go-builder /go/src/devsmake/startup.sh ./startup.sh
COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

ENV NODE_ENV production
ENV PORT 80

EXPOSE 80 8080

CMD ["sh", "startup.sh"]
