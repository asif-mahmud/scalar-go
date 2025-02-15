package loader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
)

// readJSONFileFS reads a JSON file and unmarshalls it into the provided data structure.
func readJSONFileFS[T any](iofs fs.FS, path string) (T, error) {
	var data T
	if !isJSONFile(path) {
		return data, fmt.Errorf(
			"file '%s' is not a JSON file, supported extensions are [JSON]",
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

	err = json.Unmarshal(contentBytes, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
