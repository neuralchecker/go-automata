package exporters

import (
	"fmt"
	"strings"
	"time"

	"github.com/awalterschulze/gographviz"
	"github.com/neuralchecker/go-automata/automata"
	"github.com/neuralchecker/go-automata/base_types/states"
	"github.com/neuralchecker/go-automata/interfaces"
	"golang.org/x/exp/slices"
)

type AutomatonDotExporter[T any] struct {
	// TimeStamp if true, a timestamp will be added to the name of the file.
	Timestamp bool
	FileExporter
}

func (e AutomatonDotExporter[T]) Export(automaton automata.Automaton[T], pathStr string) error {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(pathStr)
	strBuilder.WriteString("_")
	if e.Timestamp {
		strBuilder.WriteString(time.Now().UTC().Format("2006-01-02T15:04:05"))
	}
	strBuilder.WriteString(".dot")
	pathStr = strBuilder.String()

	graph, err := createGraph(automaton)
	if err != nil {
		return err
	}
	graphAst, err := graph.WriteAst()
	if err != nil {
		return err
	}
	return e.ExportToFile(pathStr, []byte(graphAst.String()))

}

func createGraph[T any](automaton automata.Automaton[T]) (*gographviz.Graph, error) {
	gName := strings.ReplaceAll(strings.ReplaceAll(automaton.GetName(), " ", "_"), "'", "")
	graph := gographviz.NewGraph()
	graph.SetDir(true)
	graph.SetName(gName)
	graph.Attrs.Ammend(gographviz.Attrs{gographviz.RankDir: "LR", gographviz.Size: "8.5"})

	// Add start node
	if err := graph.AddNode(gName, "start", map[string]string{string(gographviz.Shape): "point"}); err != nil {
		return nil, err
	}

	// Add states
	for _, state := range automaton.GetStates() {
		attrs := make(map[string]string)
		if state.IsFinal() {
			attrs[string(gographviz.Shape)] = "doublecircle"
		} else {
			attrs[string(gographviz.Shape)] = "circle"
		}
		if err := graph.AddNode(gName, state.String(), attrs); err != nil {
			return nil, err
		}

	}
	// Add starting edges
	for _, state := range automaton.GetInitialStates() {
		if err := graph.AddEdge("start", state.String(), true, nil); err != nil {
			return nil, err
		}
	}

	for _, state := range automaton.GetStates() {
		transitions := make(map[states.State[T]][]fmt.Stringer)
		for _, transition := range state.GetTransitions() {
			transitions[transition.Snd] = append(transitions[transition.Snd], transition.Fst)
		}
		for nextState, transitions := range transitions {
			switch transitions[0].(type) {
			case interfaces.Symbol[T]:
				transitionLabel := getLabelForTransitions(transitions, automaton.GetAlphabet())
				attrs := map[string]string{string(gographviz.Label): transitionLabel}
				if err := graph.AddEdge(state.String(), nextState.String(), true, attrs); err != nil {
					return nil, err
				}
			case interfaces.Guard[T]:
				for _, transition := range transitions {
					guard := transition.(interfaces.Guard[T])
					attrs := map[string]string{string(gographviz.Label): guard.String()}
					if err := graph.AddEdge(state.String(), nextState.String(), true, attrs); err != nil {
						return nil, err
					}
				}
			default:
				return nil, fmt.Errorf("unknown transition type")

			}
		}
	}

	return graph, nil
}

func getLabelForTransitions[T any](transitions []fmt.Stringer, alphabet interfaces.Alphabet[T]) string {
	if len(transitions) == alphabet.Length() {
		return "Σ"
	}

	strBuilder := strings.Builder{}
	if len(transitions) <= alphabet.Length()/2 {
		for i, transition := range transitions {
			strBuilder.WriteString(transition.String())
			if i != len(transitions)-1 {
				strBuilder.WriteString(", ")
			}
		}
	} else {
		symbols := make([]fmt.Stringer, 0, alphabet.Length()-len(transitions))
		it := alphabet.Iterator()
		for it.HasNext() {
			if !slices.Contains[fmt.Stringer](transitions, it.Next()) {
				symbols = append(symbols, it.Next())
			}
		}
		strBuilder.WriteString(`"Σ - {`)
		for i, symbol := range symbols {
			strBuilder.WriteString(symbol.String())
			if i != len(symbols)-1 {
				strBuilder.WriteString(", ")
			}
		}
		strBuilder.WriteString(`}"`)
	}

	return strBuilder.String()
}
