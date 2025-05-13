package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_requests_total",
		Help: "Total number of ping requests handled",
	},
)
	

func ping(w http.ResponseWriter, req *http.Request) {
	log.Println("Got ping! Sending pong!")
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)


	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	
	log.Println("Starting on 8090, /ping is available")
	http.ListenAndServe(":8090", nil)
}
