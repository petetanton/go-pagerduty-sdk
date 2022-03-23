package pagerduty

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

func Test_ep(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		validateMethod(t, r, "GET")
		w.Write([]byte(`{"escalation_policy": {"id": "1"}}`))
	})

	resp, err := client.GetEscalationPolicy("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &model.EscalationPolicy{
		ApiObject: model.ApiObject{
			Id: "1",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// Copied tests
// ListEscalationPolicys
func TestEP_List(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"escalation_policies": [{"id": "1"}]}`))
	})

	res, err := client.ListEscalationPolicies()

	want := &model.EscalationPolicy{
		ApiObject: model.ApiObject{
			Id: "1",
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res[0])
}

// Create EscalationPolicy
func TestEP_Create(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write([]byte(`{"escalation_policy": {"id": "1","name":"foo"}}`))
	})

	input := &model.EscalationPolicy{
		Name: "foo",
	}
	res, err := client.CreateEscalationPolicy(input)

	want := &model.EscalationPolicy{
		ApiObject: model.ApiObject{Id: "1"},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

// Delete EscalationPolicy
func TestEP_Delete(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	id := "1"
	err := client.DeleteEscalationPolicy(id)
	if err != nil {
		t.Fatal(err)
	}
}

// Get EscalationPolicy
func TestEP_Get(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"escalation_policy": {"id": "1","name":"foo"}}`))
	})

	id := "1"
	res, err := client.GetEscalationPolicy(id)

	want := &model.EscalationPolicy{
		ApiObject: model.ApiObject{Id: id},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

func TestEP_Update(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, _ = w.Write([]byte(`{"escalation_policy": {"id": "1","name":"foo", "type": "escalation_policy"}}`))
	})

	t.Run("it updates without an id", func(t *testing.T) {
		_, err := client.UpdateEscalationPolicy(&model.EscalationPolicy{})
		if err == nil {
			t.Fatal("expected an err")
		}
	})
	t.Run("it updates happy path", func(t *testing.T) {
		id := "1"
		ep := &model.EscalationPolicy{
			ApiObject: model.ApiObject{Id: id, Type: "escalation_policy"},
			Name:      "foo",
		}
		res, err := client.UpdateEscalationPolicy(ep)
		if err != nil {
			t.Fatal(err)
		}

		testEqual(t, ep, res)
	})
}
