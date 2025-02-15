package loader

import (
	"fmt"
	"io"
	"io/fs"

	"gopkg.in/yaml.v3"
)

// readYamlFileFS reads a YAML file and unmarshalls it into the provided data structure.
func readYamlFileFS[T any](iofs fs.FS, path string) (T, error) {
	var data T
	if !isYamlFile(path) {
		return data, fmt.Errorf(
			"file '%s' is not a YAML file, supported extensions are [yml|yaml]",
			path,
		)
	}

	file, err := iofs.Open(path)
	if err != nil {
		return data, err
	}

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		return data, err
	}

	err = yaml.Unmarshal(contentBytes, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
