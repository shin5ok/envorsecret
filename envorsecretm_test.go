package envorsecretm_test

import (
	"os"
	"testing"

	"github.com/shin5ok/envorsecretm"
)

func TestEnvsecretm(t *testing.T) {
	ProjectId := os.Getenv("PROJECT_ID")
	m := envorsecretm.Config{ProjectId}
	if m.Get("PROJECT_ID") != ProjectId {
		t.Logf("PROJECT_ID is NOT %s", ProjectId)
	} else {
		t.Log("PROJECT_ID is got from env correctly")
	}

}
