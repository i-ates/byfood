# JS build stage (Node.js)
FROM node:16.18.1-alpine3.16 AS js_build
WORKDIR /webapp
COPY webapp/package*.json ./
RUN npm install
COPY webapp ./
RUN npm run build

# Go build stage (Golang)
FROM golang:1.20-alpine3.16 AS go_build
WORKDIR /server
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server ./
RUN go build -o /go/bin/server

# Final stage with minimal base (Alpine + MongoDB)
FROM alpine:3.16.3

# Install MongoDB binaries only, no additional dev packages
RUN apk update && apk add --no-cache mongodb mongodb-tools

# Create MongoDB data directory
RUN mkdir -p /data/db

# Copy built JS and Go artifacts from previous stages
COPY --from=js_build /webapp/build ./webapp/
COPY --from=go_build /go/bin/server ./server

# Expose ports for Go app and MongoDB
EXPOSE 8080 27017

# Run MongoDB in background and then start Go server
CMD mongod --fork --logpath /var/log/mongodb.log --dbpath /data/db && ./server
