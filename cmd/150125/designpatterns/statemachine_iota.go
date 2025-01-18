package designpatterns

type IotaState int

const (
	StateIdle IotaState = iota
	StateProcessing
	StateCompleted
)

func (s *IotaState) IotaTransition(event string) IotaState {
	switch *s {
	case StateIdle:
		if event == "start" {
			return StateProcessing
		}
	case StateProcessing:
		if event == "complete" {
			return StateCompleted
		}
	case StateCompleted:
		return StateIdle
	}
	return *s // No state change
}

func (s IotaState) String() string {
	switch s {
	case StateIdle:
		return "Idle"
	case StateProcessing:
		return "Processing"
	case StateCompleted:
		return "Completed"
	}

	return "Unkown state"
}
