##### Stage 1 - Build Go application #####
FROM golang:1.23-alpine as builder

WORKDIR /project

# Copy Go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN go build -o /project/app ./main.go


##### Stage 2 - Final #####
# Use python-alpine as the base image to support both Go and Python
FROM python:3.9-alpine

# Set working directory
WORKDIR /dist

# Copy the Go binary from the builder stage
COPY --from=builder /project/app .

# Copy the Python script and model
COPY pkgs/python/script.py /dist/pkgs/python/script.py
COPY pkgs/python/model.pkl /dist/pkgs/python/model.pkl

# Install build tools and dependencies
RUN apk add --no-cache gcc musl-dev python3-dev g++ build-base

# Install necessary Python dependencies
RUN pip install joblib pandas scikit-learn

# Run the Go application
CMD ["./app"]
