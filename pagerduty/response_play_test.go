package pagerduty

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func Test_rp(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		validateMethod(t, r, "GET")
		w.Write([]byte(`{"response_play": {"id": "1"}}`))
	})

	resp, err := client.GetResponsePlay("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &model.ResponsePlay{
		ApiObject: model.ApiObject{
			Id: "1",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// Copied tests
// ListResponsePlays
func TestRP_List(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/response_plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"response_plays": [{"id": "1"}]}`))
	})

	res, err := client.ListResponsePlays()

	want := &model.ResponsePlay{
		ApiObject: model.ApiObject{
			Id: "1",
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res[0])
}

// Create ResponsePlay
func TestRP_Create(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/response_plays", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write([]byte(`{"response_play": {"id": "1","name":"foo"}}`))
	})

	input := &model.ResponsePlay{
		Name: "foo",
	}
	res, err := client.CreateResponsePlay(input)

	want := &model.ResponsePlay{
		ApiObject: model.ApiObject{Id: "1"},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

// Delete ResponsePlay
func TestRP_Delete(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	id := "1"
	err := client.DeleteResponsePlay(id)
	if err != nil {
		t.Fatal(err)
	}
}

// Get ResponsePlay
func TestRP_Get(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"response_play": {"id": "1","name":"foo"}}`))
	})

	id := "1"
	res, err := client.GetResponsePlay(id)

	want := &model.ResponsePlay{
		ApiObject: model.ApiObject{Id: id},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

func TestRP_Update(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/response_plays/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, _ = w.Write([]byte(`{"response_play": {"id": "1","name":"foo", "type": "response_play"}}`))
	})

	t.Run("it updates without an id", func(t *testing.T) {
		_, err := client.UpdateResponsePlay(&model.ResponsePlay{})
		if err == nil {
			t.Fatal("expected an err")
		}
	})
	t.Run("it updates happy path", func(t *testing.T) {
		id := "1"
		ep := &model.ResponsePlay{
			ApiObject: model.ApiObject{Id: id, Type: "response_play"},
			Name:      "foo",
		}
		res, err := client.UpdateResponsePlay(ep)
		if err != nil {
			t.Fatal(err)
		}

		testEqual(t, ep, res)
	})
}
