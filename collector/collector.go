package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/exporter/httpclient"
)

type Collector interface {
	Describe(ch chan<- *prometheus.Desc)
	Collect(ch chan<- prometheus.Metric)
}

func NewCollector(urls []string, client httpclient.Client) Collector {
	return newCollector(urls, client)
}

type collector struct {
	statusMetric   *prometheus.Desc
	responseMetric *prometheus.Desc
	url            []string
	httpclient     httpclient.Client
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newCollector(urls []string, client httpclient.Client) *collector {
	return &collector{
		statusMetric: prometheus.NewDesc("sample_external_url_up",
			"",
			[]string{"url"}, nil,
		),
		responseMetric: prometheus.NewDesc("sample_external_url_response_ms",
			"",
			[]string{"url"}, nil,
		),
		url:        urls,
		httpclient: client,
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *collector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.statusMetric
	ch <- collector.responseMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *collector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	for _, url := range collector.url {
		status, response := collector.httpclient.Get(url)

		//Write latest value for each metric in the prometheus metric channel.
		//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
		ch <- prometheus.MustNewConstMetric(collector.statusMetric, prometheus.CounterValue, status, url)
		ch <- prometheus.MustNewConstMetric(collector.responseMetric, prometheus.CounterValue, response, url)
	}
}
