package definitions_test

import (
	"testing"

	"github.com/neuralchecker/go-automata/automata/definitions"
	"github.com/neuralchecker/go-automata/exporters"
	"github.com/stretchr/testify/assert"
)

func TestTomitas1(t *testing.T) {
	dfa := definitions.GetTomitas1()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}

func TestTomitas2(t *testing.T) {
	dfa := definitions.GetTomitas2()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}

func TestTomitas3(t *testing.T) {
	dfa := definitions.GetTomitas3()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}

func TestTomitas4(t *testing.T) {
	dfa := definitions.GetTomitas4()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}

func TestTomitas5(t *testing.T) {
	dfa := definitions.GetTomitas5()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}

func TestTomitas6(t *testing.T) {
	dfa := definitions.GetTomitas6()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}

func TestTomitas7(t *testing.T) {
	dfa := definitions.GetTomitas7()
	assert.True(t, dfa.IsDeterministic())
	dfa.AddExporter(exporters.AutomatonDotExporter[rune]{Timestamp: true})
	err := dfa.Export("./testdata")
	assert.NoError(t, err)
}
