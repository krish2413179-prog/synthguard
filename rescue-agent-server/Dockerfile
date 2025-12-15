# 1. Start with a lightweight Go image
FROM golang:1.25-alpine

# 2. Set the folder inside the container
WORKDIR /app

# 3. Copy only the dependency files first (better caching)
COPY go.mod go.sum ./

# 4. Download the dependencies
RUN go mod download

# 5. Copy the rest of your code
COPY . .

# 6. Build the Go app into a binary named "agent"
RUN go build -o agent main.go

# 7. Start the app
CMD ["./agent"]