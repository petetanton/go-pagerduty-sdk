package pagerduty

import (
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_errors(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		validateMethod(t, r, "POST")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": {"message": "Invalid Input Provided","code": 2001,"errors": ["Name must be a String."]}}`))
	})

	_, err := client.CreateUser(&model.User{
		Name: "my-cool-user",
	})
	assert.Error(t, err, "an error is required")
	assert.Equal(t, err.Error(), "got a 400 response from PagerDuty: Name must be a String.")
}
