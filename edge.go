package argon

type Edge struct {
	From     State
	To       State
	Action   string
	Callback bool
}
