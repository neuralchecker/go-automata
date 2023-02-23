package alphabets

import "github.com/neuralchecker/go-automata/interfaces"

func Join[T any](alphabets ...interfaces.Alphabet[T]) interfaces.Alphabet[T] {
	joinedSymbols := make(map[interfaces.Symbol[T]]struct{})
	for _, alphabet := range alphabets {
		it := alphabet.Iterator()
		for it.HasNext() {
			joinedSymbols[it.Next()] = struct{}{}
		}
	}
	symbols := make([]interfaces.Symbol[T], 0, len(joinedSymbols))
	for symbol := range joinedSymbols {
		symbols = append(symbols, symbol)
	}
	return New(symbols)
}
