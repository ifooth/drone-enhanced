FROM golang:1.15.2-buster as builder

ARG PROMU_VERSION=0.12.0

RUN wget -q https://github.com/prometheus/promu/releases/download/v${PROMU_VERSION}/promu-${PROMU_VERSION}.linux-amd64.tar.gz && \
    tar -xf promu-${PROMU_VERSION}.linux-amd64.tar.gz && \
    mv promu-${PROMU_VERSION}.linux-amd64/promu bin

WORKDIR /app

# COPY go.mod /app
# COPY go.sum /app
# RUN go mod download

COPY . /app

RUN promu build

FROM debian:buster

WORKDIR /app
COPY --from=builder /app/drone_enhanced /usr/local/bin/drone_enhanced

CMD ["/usr/local/bin/drone_enhanced", "server", "--http-address", "0.0.0.0:8080"]
