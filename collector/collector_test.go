package collector

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/exporter/httpclient"
	fake "github.com/exporter/mocks/httpclient"

	"github.com/prometheus/client_golang/prometheus"
)

func TestCollector(t *testing.T) {

	client := fake.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	c := httpclient.NewClient(client)

	collect := NewCollector([]string{"someurl"}, c)
	metricChan := make(chan prometheus.Metric)
	go func() {
		collect.Collect(metricChan)
		close(metricChan)
	}()
	for m := range metricChan {
		if strings.Contains(m.Desc().String(), "sample_external_url_up") || strings.Contains(m.Desc().String(), "sample_external_url_response_ms") {
			continue
		} else {
			t.Fatalf(m.Desc().String())
		}

	}

}
