FROM golang:1.21

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o cmd cmd/main.go

CMD ["cmd/main"]
