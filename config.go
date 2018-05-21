package argon

type Config struct {
	StartState State
	States     []State
	Edges      []Edge
}
