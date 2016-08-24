package main

import (
	"github.com/GeertJohan/go-metrics/influxdb"
	"github.com/rcrowley/go-metrics"
	"net/http"
	"time"
)

var requestCounter metrics.Counter
var responseTime metrics.Timer

func MetricToInfluxDB(d time.Duration) {
	go influxdb.Influxdb(metrics.DefaultRegistry, d, &influxdb.Config{
		Host:     "localhost:8086",
		Database: "metric",
		Username: "root",
		Password: "root",
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc(1)
	startReqTime := time.Now()
	defer responseTime.Update(time.Since(startReqTime))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func main() {
	requestCounter = metrics.NewCounter()
	metrics.Register("count_request", requestCounter)

	responseTime = metrics.NewTimer()
	metrics.Register("response_time", responseTime)

	MetricToInfluxDB(time.Second * 1)

	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":3000", nil)
}
