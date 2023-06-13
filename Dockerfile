FROM golang:1.20.5-bullseye as builder
WORKDIR /app
COPY . /app
RUN go build -o http-server

FROM debian:11.2-slim as final
WORKDIR /opt/http
COPY . /app
COPY --from=builder /app/http-server /opt/http
CMD ["/opt/http/http-server"]
