package main

import "github.com/prometheus/client_golang/prometheus"

var (
	imgServed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "xtradio_cdn_image_served",
			Help: "Number of images served.",
		},
	)
	// hdFailures = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "hd_errors_total",
	// 		Help: "Number of hard-disk errors.",
	// 	},
	// 	[]string{"device"},
	// )
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(imgServed)
	// prometheus.MustRegister(hdFailures)
}
