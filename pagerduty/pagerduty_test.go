package pagerduty

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux *http.ServeMux

	server *httptest.Server
)

func setup() *Client {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, err := NewClient(&Config{ApiUrl: server.URL, ApiToken: "foo"})
	if err != nil {
		panic(err)
	}
	return client
}

func teardown() {
	server.Close()
}

func validateMethod(t *testing.T, r *http.Request, method string) {
	assert.Equalf(t, method, r.Method, "expected a %s request by got a %s", method, r.Method)
}
