FROM golang:1.18 as builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

ENV GOPROXY https://proxy.golang.org,direct
ENV CGO_ENABLED=0

RUN go mod download

COPY . .

RUN GOOS=linux go build -o auth


##### Stage 2 #####

### Define the running image
FROM scratch

### Alternatively to 'FROM scratch', use 'alpine':
# FROM alpine:3.13.1

### Set working directory
WORKDIR /app

### Copy built binary application from 'builder' image
COPY --from=builder /app .

### Run the binary application
CMD ["/app/auth"]