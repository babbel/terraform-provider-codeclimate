package codeclimateclient

import (
	"fmt"
	"net/http"
	"testing"
)

const (
	repositoryId           = "5b6abdc65b6abdc65b6abdc6"
	expectedTestReporterId = "0c89092bc2c088d667612ddd1a992ec62f643ded331f40783bcf6b847561234d"
)

func TestGetId(t *testing.T) {
	teardown := setup()
	defer teardown()

	handURL := fmt.Sprintf("/repos/%s", repositoryId)

	mux.HandleFunc(handURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getFixture("repositories/repository.json"))
	})

	repository, err := client.GetRepository(repositoryId)

	if err != nil {
		t.Fatal(err)
	}

	if repository.TestReporterId != expectedTestReporterId {
		t.Errorf("Expected test_reporter_id to be '%s', got: '%s'", expectedTestReporterId, repository.TestReporterId)
	}
}
