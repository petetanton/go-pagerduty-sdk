package pagerduty

import (
	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
	"net/http"
	"reflect"
	"testing"
)

func Test_team(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		validateMethod(t, r, "GET")
		w.Write([]byte(`{"team": {"id": "1"}}`))
	})

	resp, err := client.GetTeam("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &model.Team{
		ApiObject{Id: "1"},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}
