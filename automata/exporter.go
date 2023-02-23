package automata

import (
	"path"
	"strings"
)

type Exporter[T any] interface {
	// Export exports the model to the specified folder.
	Export(model Automaton[T], path string) error
}

// ResolvePath resolves the path of the exported model. If the path is an empty string, it will be resolved to ./output_models.
// You should still add the extension to the path, as it will not be added automatically.
//
// Parameters:
//
//	model: the model to be exported.
//	pathStr: the path to be resolved. If you do not want to use output_models, you should pass ".".
//	timestamp: if true, a timestamp will be added to the name of the file.
//
// Returns:
//
//	the resolved path
func ResolvePath[T any](model Automaton[T], pathStr string) string {
	if pathStr == "" {
		pathStr = "./output_models"
	}
	nameBuilder := strings.Builder{}
	nameBuilder.WriteString(model.GetName())

	pathStr = path.Join(pathStr, nameBuilder.String())
	return pathStr
}
