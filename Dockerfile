FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./application

FROM alpine:latest
COPY --from=builder /build .
EXPOSE 4444
ENTRYPOINT ["./main"]