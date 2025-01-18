package designpatterns

import "fmt"

type FuncState func(event string) FuncState

func StateIdleF(event string) FuncState {
	if event == "start" {
		fmt.Println("transitioning from Idle to Processing")
		return StateProcessingF
	}
	return StateIdleF
}

func StateProcessingF(event string) FuncState {
	if event == "complete" {
		fmt.Println("transitioning from Processing to Completed")
		return StateCompletedF
	}
	return StateProcessingF
}

func StateCompletedF(_ string) FuncState {
	fmt.Println("returning to idle state")
	return StateIdleF
}
