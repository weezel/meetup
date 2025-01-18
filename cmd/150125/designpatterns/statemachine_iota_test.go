package designpatterns

import (
	"log"
	"testing"
)

func TestState_IotaTransition(t *testing.T) {
	t.Helper()

	// Only demonstrates how the state transitions happen.
	// Writing tests is left as an exercise for the reader.
	sm := StateIdle
	log.Printf("Initial state: %d\n", sm)
	events := []string{"start", "complete", "reset"}
	for _, event := range events {
		newState := sm.IotaTransition(event)
		log.Printf("Event: %s, Transitioned from %s (iota=%d) to %s (iota=%d)\n",
			event,
			sm,
			sm,
			newState,
			newState,
		)
		sm = newState
	}
}
