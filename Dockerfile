FROM golang:1.16-alpine AS base

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go get github.com/Looty/cli-test/cmd
RUN go build -o /cli-test

# Start fresh from a smaller image
FROM alpine:3.9

COPY --from=base /app/cli-test /app/cli-test

# Run the binary program produced by `go install`
CMD ["/app/cli-test"]