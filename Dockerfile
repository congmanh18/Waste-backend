##### Stage 1 #####
FROM golang:1.23-alpine as builder


RUN mkdir -p /project
WORKDIR /project

### Copy Go application dependency files
COPY go.mod .
COPY go.sum .

### Download Go application module dependencies
RUN go mod download


### Copy actual source code for building the application
COPY . .

ENV CGO_ENABLED=0

RUN go build -o app main.go


##### Stage 2 #####
FROM scratch

WORKDIR /dist 

### Copy the .env file
COPY --from=builder /project/app .
COPY --from=builder /project/.env .

CMD ["./app"]