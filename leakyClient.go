package leakyClient

import (
	"io"
	"net/http"
	"net/url"

	"go.uber.org/ratelimit"
)

// LeakyClient Provides a leaky-bucket rate-limited HTTP client
type LeakyClient struct {
	limiter ratelimit.Limiter
	inner   http.Client
}

// New Constructs a new LeakyClient that handles a maximum of `rate` requests per second,
// and an optional inner http.Client (if client is not provided, the default http.Client is used).
func New(rate int, client ...http.Client) LeakyClient {
	if len(client) > 0 {
		return LeakyClient{
			ratelimit.New(rate),
			client[0],
		}
	} else {
		return LeakyClient{
			ratelimit.New(rate),
			http.Client{},
		}
	}
}

// Do Performs an HTTP request using the inner HTTP client.
func (lc LeakyClient) Do(req *http.Request) (*http.Response, error) {
	lc.limiter.Take()
	return lc.inner.Do(req)
}

// Get Issues a GET to the specified URL.
func (lc LeakyClient) Get(url string) (*http.Response, error) {
	lc.limiter.Take()
	return lc.inner.Get(url)
}

// Head Issues a HEAD to the specified URL.
func (lc LeakyClient) Head(url string) (*http.Response, error) {
	lc.limiter.Take()
	return lc.inner.Head(url)
}

// Post Issues a POST to the specified URL.
func (lc LeakyClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	lc.limiter.Take()
	return lc.inner.Post(url, contentType, body)
}

// PostForm Issues a POST to the specified URL with the specified data.
func (lc LeakyClient) PostForm(url string, data url.Values) (*http.Response, error) {
	lc.limiter.Take()
	return lc.inner.PostForm(url, data)
}

// CloseIdleConnections closes any connections on its Transport which were previously connected from previous requests but are now sitting idle in a "keep-alive" state.
// It does not interrupt any connections currently in use.
func (lc LeakyClient) CloseIdleConnections() {
	lc.inner.CloseIdleConnections()
}
