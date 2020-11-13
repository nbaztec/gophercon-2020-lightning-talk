package main

import (
	"functional-opts/metrics"
)

const (
	LatencyLow  = 10
	LatencyMid  = 100
	LatencyHigh = 1000
)

func main() {

	// count a http request
	metrics.Count("http_request", metrics.Category("area-51"), metrics.Latency(LatencyHigh))

	// count a db error
	metrics.Count("db_error", metrics.Tags([]string{"pg", "write"}))
}
