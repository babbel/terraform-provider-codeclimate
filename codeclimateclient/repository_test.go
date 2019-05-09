package codeclimateclient

import (
	"fmt"
	"net/http"
	"testing"
)

const (
	repositorySlug         = "lessonnine/testarepo"
	expectedTestReporterId = "0c89092bc2c088d667612ddd1a992ec62f643ded331f40783bcf6b847561234d"
)

func TestGetId(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("github_slug") != repositorySlug {
			t.Fatal(fmt.Errorf("received slug doesn match `%s`", repositorySlug))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/repository.json"))
	})

	repository, err := client.GetRepository(repositorySlug)

	if err != nil {
		t.Fatal(err)
	}

	if repository.TestReporterId != expectedTestReporterId {
		t.Errorf("Expected test_reporter_id to be '%s', got: '%s'", expectedTestReporterId, repository.TestReporterId)
	}
}
