FROM golang:1.22-bullseye as builder
WORKDIR /go/src/lambda
COPY go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN GOARCH=amd64 GOOS=linux go build -a -o /main ./cmd/example

FROM debian:bullseye as runner
WORKDIR /app/
COPY --from=builder /main /main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "/main" ]
