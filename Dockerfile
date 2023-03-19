FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o auth

EXPOSE 5000

CMD ./auth