package statemachine

type EventListener interface {
	OnStart()
	OnStop()
}
