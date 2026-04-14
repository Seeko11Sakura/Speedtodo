package focus

import "testing"

func TestCompleteSessionChangesStatus(t *testing.T) {
	session := Session{Status: "running"}
	session.Complete()
	if session.Status != "completed" {
		t.Fatal("expected status to become completed")
	}
}
