package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	registry := prometheus.NewRegistry()
	appHealth := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "test",
		Name:      "app_health",
		Help:      "app health, 1 for up, 0 for down",
	})
	go func() {
		for {
			fakeHealth := float64(rand.Intn(2))
			appHealth.Set(fakeHealth)
			time.Sleep(time.Second)
		}
	}()
	registry.MustRegister(appHealth)
	fmt.Println("Wish you good day ^^")
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(":9101", nil)
}
