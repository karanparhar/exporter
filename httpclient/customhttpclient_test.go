package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	fake "github.com/exporter/mocks/httpclient"
)

func TestGetWithRoundTripper(t *testing.T) {

	client := fake.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	api := httpClient{client}
	status, _ := api.Get("http://example.com")
	assert.Equal(t, status, float64(1))

}

func TestGetWithRoundTripperFailed(t *testing.T) {

	client := fake.NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 503,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`NOT OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	api := httpClient{client}
	status, _ := api.Get("http://example.com")
	assert.Equal(t, status, float64(0))

}
