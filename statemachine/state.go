package statemachine

type State interface {
	EventListener

	Entry()
	Do()
	Exit()
}
