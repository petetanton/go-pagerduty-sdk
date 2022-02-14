package pagerduty

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"text/template"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
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
		ApiObject: model.ApiObject{
			Id: "1",
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

// Copied tests
// ListTeams
func TestTeam_List(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"teams": [{"id": "1"}]}`))
	})

	res, err := client.ListTeams()

	want := &model.Team{
		ApiObject: model.ApiObject{
			Id: "1",
		},
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res[0])
}

// Create Team
func TestTeam_Create(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = w.Write([]byte(`{"team": {"id": "1","name":"foo"}}`))
	})

	input := &model.Team{
		Name: "foo",
	}
	res, err := client.CreateTeam(input)

	want := &model.Team{
		ApiObject: model.ApiObject{Id: "1"},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

// Delete Team
func TestTeam_Delete(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	id := "1"
	err := client.DeleteTeam(id)
	if err != nil {
		t.Fatal(err)
	}
}

// Get Team
func TestTeam_Get(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = w.Write([]byte(`{"team": {"id": "1","name":"foo"}}`))
	})

	id := "1"
	res, err := client.GetTeam(id)

	want := &model.Team{
		ApiObject: model.ApiObject{Id: id},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

// Update Team
func TestTeam_Update(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, _ = w.Write([]byte(`{"team": {"id": "1","name":"foo"}}`))
	})

	id := "1"

	input := &model.Team{
		ApiObject: model.ApiObject{Id: id},
		Name:      "foo",
	}
	res, err := client.UpdateTeam(input)

	want := &model.Team{
		ApiObject: model.ApiObject{Id: id},
		Name:      "foo",
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, want, res)
}

// Remove Escalation Policy from Team
//func TestTeam_RemoveEscalationPolicyFromTeam(t *testing.T) {
//	setup()
//	defer teardown()
//
//	mux.HandleFunc("/teams/1/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, "DELETE")
//	})
//
//	client := setup()
//	teamID := "1"
//	epID := "1"
//
//	err := client.RemoveEscalationPolicyFromTeam(teamID, epID)
//	if err != nil {
//		t.Fatal(err)
//	}
//}

//Add Escalation Policy to Team
func TestTeam_AddEscalationPolicyToTeam(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1/escalation_policies/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	teamID := "1"
	epID := "1"

	err := client.AddEscalationPolicyToTeam(teamID, epID)
	if err != nil {
		t.Fatal(err)
	}
}

// Remove User from Team
func TestTeam_RemoveUserFromTeam(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	teamID := "1"
	userID := "1"

	err := client.RemoveUserFromTeam(teamID, userID)
	if err != nil {
		t.Fatal(err)
	}
}

// Add User to Team
func TestTeam_AddUserToTeam(t *testing.T) {
	client := setup()
	defer teardown()

	mux.HandleFunc("/teams/1/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusNoContent)
	})

	teamID := "1"
	userID := "1"

	err := client.AddUserToTeam(teamID, userID, "some_role")
	if err != nil {
		t.Fatal(err)
	}
}

func userID(offset, index int) int {
	return offset + index
}

func lastIndex(length, index int) bool {
	return length-1 == index
}

const membersResponseTemplate = `
{
    {{template "pageInfo" . }},
    "members": [
        {{- $length := len .roles -}}
        {{- $offset := .offset -}}
        {{- range $index, $role := .roles -}}
        {
            "user": {
                "id": "ID{{userID $offset $index}}"
            },
            "role": "{{ $role }}"
        }
        {{- if not (lastIndex $length $index) }},
        {{end -}}
        {{- end }}
    ]
}
`

var memberPageTemplate = template.Must(pageTemplate.New("membersResponse").
	Funcs(templateUtilityFuncs).
	Parse(membersResponseTemplate))

const (
	testValidTeamID = "MYTEAM"
	testAPIKey      = "MYKEY"
	testBadURL      = "A-FAKE-URL"
	testMaxPageSize = 3
)

var templateUtilityFuncs = template.FuncMap{
	"lastIndex": lastIndex,
	"userID":    userID,
}

var pageTemplate = template.Must(template.New("pageInfo").Parse(`
    "more": {{- .more -}},
    "limit": {{- .limit -}},
    "offset": {{- .offset -}}
`))

type pageDetails struct {
	lowNumber, highNumber, limit, offset int
	more                                 bool
}

func genMembersRespPage(details pageDetails, t *testing.T) string {
	if details.limit == 0 {
		details.limit = 25 // Default to 25, PD's API default.
	}

	possibleRoles := []string{"manager", "responder", "observer"}
	roles := make([]string, 0)
	for ; details.lowNumber <= details.highNumber; details.lowNumber++ {
		roles = append(roles, possibleRoles[details.lowNumber%len(possibleRoles)])
	}

	buffer := bytes.NewBufferString("")
	err := memberPageTemplate.Execute(buffer, map[string]interface{}{
		"roles":  roles,
		"more":   details.more,
		"limit":  details.limit,
		"offset": details.offset,
	})
	if err != nil {
		t.Fatalf("Failed to apply values to template: %v", err)
	}

	return string(buffer.String())
}

func genRespPages(amount,
	maxPageSize int,
	pageGenerator func(pageDetails, *testing.T) string,
	t *testing.T) []string {
	pages := make([]string, 0)

	lowNumber := 1
	offset := 0
	more := true

	for {
		tempHighNumber := amount

		if lowNumber+(maxPageSize-1) < amount {
			// Still more pages to come, this page doesn't hit upper.
			tempHighNumber = lowNumber + (maxPageSize - 1)
		} else {
			// Last page, with at least 1 user.
			more = false
		}

		// Generate page using current lower and upper.
		page := pageGenerator(pageDetails{
			lowNumber:  lowNumber,
			highNumber: tempHighNumber,
			limit:      maxPageSize,
			more:       more,
			offset:     offset,
		}, t)

		pages = append(pages, page)

		if !more {
			// Hit the last page, stop.
			return pages
		}
		// Move the offset and lower up to prepare for next page.
		offset += maxPageSize
		lowNumber += maxPageSize
	}
}

func TestListMembersSuccess(t *testing.T) {
	client := setup()
	defer teardown()

	expectedNumResults := testMaxPageSize - 1
	page := genRespPages(expectedNumResults, testMaxPageSize, genMembersRespPage, t)[0]

	mux.HandleFunc("/teams/"+testValidTeamID+"/members", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
	})

	members, err := client.ListTeamMembers(testValidTeamID)
	if err != nil {
		t.Fatalf("Failed to get members: %v", err)
	}

	if len(members) != expectedNumResults {
		t.Fatalf("Expected %d team members, got: %v", expectedNumResults, err)
	}
}

func TestListMembersError(t *testing.T) {
	client := setup()
	members, err := client.ListTeamMembers(testValidTeamID)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if members != nil {
		t.Fatalf("Expected nil members response, got: %v", members)
	}
}
