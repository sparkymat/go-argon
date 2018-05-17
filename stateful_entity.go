package argon

type State string

type StatefulEntity interface {
	GetState() State
	SetState(newState State)
}
