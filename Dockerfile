FROM golang:1.15.7 as builder
WORKDIR /sensors_collector

RUN \
    go get github.com/lib/pq
COPY . .
RUN go build .

FROM golang:1.15.7
VOLUME /conf
ENTRYPOINT ["/go/bin/sensors_collector", "-dbConfigPath", "/conf/db_config.json", "-sensorsConfigPath", "/conf/sensors_config.json"]

COPY --from=builder /sensors_collector/sensors_collector /go/bin/
