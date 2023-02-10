package automata

type Exporter[T any] interface {
	// Export exports the model to the specified folder.
	Export(model FiniteAutomaton[T], path string)
}
