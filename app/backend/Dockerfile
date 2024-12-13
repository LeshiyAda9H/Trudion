FROM golang:latest

WORKDIR /app/backend

COPY . .
RUN go mod download

EXPOSE 8080

RUN go build -o /godocker
