package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	c "github.com/exporter/collector"
	cf "github.com/exporter/config"
	"github.com/exporter/httpclient"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var conf = flag.String("c", "", "config file to read url endpoints")

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		log.Fatal("please provide config file")
	}

	config, err := cf.ReadConfigFile(*conf)
	if err != nil {
		log.Fatal(err)

	}

	//Create a new instance of the collector and
	//register it with the prometheus client.
	collector := c.NewCollector(config.Urls, httpclient.NewClient(&http.Client{}))
	prometheus.MustRegister(collector)
	prometheus.Unregister(collectors.NewGoCollector())

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	//healthcheck
	http.HandleFunc("/healthcheck", Self)
	log.Info("Beginning to serve on port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func Self(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "I am alive\n")
}
