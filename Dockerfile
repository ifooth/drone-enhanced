FROM golang:1.15 as builder

WORKDIR /app

COPY . /app

RUN make build

FROM alpine:3.14

COPY --from=builder /app/drone-enhanced /bin/drone-enhanced

CMD ["/bin/drone-enhanced", "server", "--http-address", "0.0.0.0:8080"]
