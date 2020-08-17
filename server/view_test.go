package server

import (
	"strings"
	"testing"
)

func TestViewIsHandledCorrectly(t *testing.T) {
	s := createDummyServer()
	output, _, err := s.view("viewThatDoesntExist")
	if err == nil {
		t.Error("We should have recieved an error because that view file shouldn't exist")
	}

	if !strings.Contains(string(output), "Totes bummer! Go check the logs to see what went wrong.") {
		t.Error("The expected error message was not returned. Output:\n" + string(output))
	}
}
