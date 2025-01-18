package designpatterns

import (
	"fmt"
	"testing"
)

func TestFuncStates(t *testing.T) {
	t.Helper()

	// Only demonstrates how the state transitions happen.
	// Writing tests is left as an exercise for the reader.
	currentState := StateIdleF
	events := []string{"start", "complete", "reset"}
	for _, event := range events {
		fmt.Printf("Event: %s, ", event)
		currentState = currentState(event)
	}
}
