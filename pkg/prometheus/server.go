package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func Serve(listen string) error {
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting listening on %v", listen)
	return http.ListenAndServe(listen, nil)
}
