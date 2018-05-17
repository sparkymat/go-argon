package argon

type Edge struct {
	From          State
	To            State
	Action        string
	OnCallback    bool
	AfterCallback bool
}
