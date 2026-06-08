FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main

EXPOSE 8080

CMD ["./main"]