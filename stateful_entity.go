package argon

type State string

type StatefulEntity interface {
	GetState() State
	SetState(newState State)
	BeforeAction(action string)
	AfterAction(action string, err error)
	OnAction(action string) error
}
