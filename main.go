package main

import "study-statemachine/statemachine"

func main() {
	machine := statemachine.Start()

	machine.OnStart()
	machine.OnStart()
	machine.OnStart()

	machine.OnStop()
}
