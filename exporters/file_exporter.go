package exporters

import (
	"os"
	"path"
)

type FileExporter struct{}

func (f FileExporter) ExportToFile(pathStr string, content []byte) error {
	dir := path.Dir(pathStr)
	os.MkdirAll(dir, os.ModePerm)

	return os.WriteFile(pathStr, content, 0644)
}
