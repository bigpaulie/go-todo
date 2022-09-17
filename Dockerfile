FROM golang:1.19-alpine

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./cmd/web/main.go && chmod +x /usr/local/bin/app

EXPOSE 8080
CMD app