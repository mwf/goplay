package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	v := prometheus.NewHistogramVec(prometheus.HistogramOpts{}, []string{"my_label"})

	var hist prometheus.Histogram
	hist = v.WithLabelValues("ololo")
	fmt.Printf("%T\n", hist)
}
