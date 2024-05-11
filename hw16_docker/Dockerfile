FROM golang:latest AS builder
RUN git clone -b hw16_docker https://github.com/TOsmanov/my-hw.git /usr/local/go/src/go_sql/
WORKDIR /usr/local/go/src/go_sql
RUN go clean --modcache && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
-mod=readonly -o go_sql /usr/local/go/src/go_sql/hw15_go_sql/server/main.go
FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/go/src/go_sql/go_sql /app/
CMD ["/app/go_sql"]
