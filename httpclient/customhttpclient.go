package httpclient

import (
	"net/http"
	"time"
)

type Client interface {
	Get(url string) (float64, float64)
}

type httpClient struct {
	Client *http.Client
}

func NewClient(c *http.Client) Client {
	return &httpClient{c}
}

func (c *httpClient) Get(url string) (float64, float64) {
	start := time.Now()

	result, err := c.Client.Get(url)

	if err != nil {
		return 0, time.Since(start).Seconds()
	}

	defer result.Body.Close()

	if result.StatusCode == http.StatusOK {

		return 1, time.Since(start).Seconds()
	}

	return 0, float64(time.Since(start).Milliseconds())
}
