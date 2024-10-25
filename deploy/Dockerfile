FROM golang:latest

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

RUN go mod download

# Set WORKDIR to cmd
WORKDIR /app/cmd

# List current directory contents
RUN ls -R .

# Copy the rest of the application
COPY . .

# Build and execute the Go binary
RUN ls -R ./cmd/app/main.go && go build -o ./cmd/main ./cmd/app/main.go && ls -R ./cmd/main && chmod +x ./cmd/main || echo "Error building main.go"

# Verify the contents of the cmd directory
RUN ls -R /app/cmd

EXPOSE 8080

CMD ["./cmd/main"]
