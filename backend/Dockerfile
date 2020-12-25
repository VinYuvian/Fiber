FROM golang:1.14
WORKDIR /app
RUN apt update && apt install -y git
COPY go.mod go.sum ./
RUN go mod download



