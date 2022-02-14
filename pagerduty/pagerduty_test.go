package pagerduty

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()

	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("values not equal (-want / +got):\n%s", diff)
	}
}
