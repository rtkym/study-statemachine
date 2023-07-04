package statemachine

import "fmt"

func New() *MachineContext {
	mc := &MachineContext{}
	state := LightsOut{mc: mc}
	mc.current = &state

	return mc
}

func Start() *MachineContext {
	m := New()
	m.current.Entry()

	return m
}

type MachineContext struct {
	current State
}

func (m *MachineContext) changeState(newState State) {
	older := m.current
	m.current = newState

	if older != nil {
		older.Exit()
	}

	m.current.Entry()
}

func (m *MachineContext) terminate() {
	fmt.Println("Terminate")
}

var _ EventListener = (*MachineContext)(nil)

func (m *MachineContext) OnStart() {
	m.current.OnStart()
}

func (m MachineContext) OnStop() {
	m.current.OnStop()
}
