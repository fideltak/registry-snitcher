package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	"github.com/fideltak/registry-snicher/internal/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type PromServer struct {
	Address string `env:"RS_PROM_ADDRESS" envDefault:"0.0.0.0"`
	PORT    string `env:"RS_PROM_PORT" envDefault:"9100"`
}

var (
	version      = "Development"
	opsSucceeded = promauto.NewCounter(prometheus.CounterOpts{
		Name: "snitch_registry_succeeded_total",
		Help: "The total number of succeeded access",
	})
	opsFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "snitch_registry_failed_total",
		Help: "The total number of failed access",
	})
)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	debugFlag := os.Getenv("RS_DEBUG")
	if debugFlag != "" {
		log.SetLevel(log.DebugLevel)

	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	log.Infof("Start Registry Snitcher: %s", version)

	log.Info("Start Prometheus Exporter")
	prom := &PromServer{}
	env.Parse(prom)
	go func() {
		log.Infof("Prometheus Exporter Address: %s:%s", prom.Address, prom.PORT)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf("%s:%s", prom.Address, prom.PORT), nil)
	}()

	testFlag := os.Getenv("RS_TEST")
	log.Debugf("Test Flag: %s", testFlag)

	defaultDuration := 60
	dStr := os.Getenv("RS_DURATION_SEC")
	log.Debugf("Duration: %s", dStr)
	d, err := strconv.Atoi(dStr)
	if err != nil {
		log.Warn(err)
		log.Infof("Set default duration sec: 60s")
		d = defaultDuration
	}
	if d == 0 {
		log.Infof("Set default duration sec: 60s")
		d = defaultDuration
	}

	image := &app.ContainerImage{}
	env.Parse(image)
	for {
		err = image.Inspect()
		if err != nil {
			log.Error(err)
			log.Debugf("Register an access count as failed to Prometheus Exporter")
			opsFailed.Inc()
		} else {
			log.Infof("Succeeded: %s", image.ImageName)
			log.Debugf("Register an access count as succeeded to Prometheus Exporter")
			opsSucceeded.Inc()
		}
		if testFlag != "" {
			break
		}
		time.Sleep(time.Duration(d) * time.Second)
		log.Debugf("Wait for %d sec", d)
	}
}
