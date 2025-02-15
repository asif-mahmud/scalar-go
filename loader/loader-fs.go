package loader

import (
	"fmt"
	"io/fs"
	"maps"

	"github.com/bdpiprava/scalar-go/model"
	"github.com/bdpiprava/scalar-go/sanitizer"
)

// LoadWithNameFS reads the API specification from the provided file system handle rootFS.
func LoadWithNameFS(rootFS fs.FS, apiFileName string) (*model.Spec, error) {
	content, err := readFileFS[model.Spec](rootFS, apiFileName)
	if err != nil {
		return nil, err
	}

	specContent := &content
	specContent.Paths = initializeIfNil(specContent.Paths)
	specContent.Components.Schemas = initializeIfNil(specContent.Components.Schemas)
	specContent.Components.Parameters = initializeIfNil(specContent.Components.Parameters)
	specContent.Components.Responses = initializeIfNil(specContent.Components.Responses)

	paths, err := readDirRecursivelyFS(rootFS, "paths")
	if err != nil {
		return nil, err
	}
	maps.Copy(specContent.Paths, *paths)

	responses, err := readDirRecursivelyFS(rootFS, "responses")
	if err != nil {
		return nil, err
	}
	maps.Copy(specContent.Components.Responses, *responses)

	schemas, err := readDirRecursivelyFS(rootFS, "schemas")
	if err != nil {
		return nil, err
	}
	maps.Copy(specContent.Components.Schemas, *schemas)

	return sanitizer.Sanitize(specContent), nil
}

// readFileFS reads a file and unmarshalls it into the provided data structure.
func readFileFS[T any](iofs fs.FS, path string) (data T, err error) {
	if data, err = readYamlFileFS[T](iofs, path); err == nil {
		return
	} else if data, err = readJSONFileFS[T](iofs, path); err == nil {
		return
	}
	return data, fmt.Errorf(
		"file '%s' is not a YAML or JSON file, supported extensions are [yml|yaml|json]",
		path,
	)
}
