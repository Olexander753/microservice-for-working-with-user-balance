FROM golang:1.19-alpine3.16
RUN mkdir app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o main ./cmd/main
CMD sleep 10 && ./main