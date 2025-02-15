package loader

import (
	"io/fs"
	"maps"
	"path/filepath"
	"strings"

	"github.com/bdpiprava/scalar-go/model"
)

// readDirRecursivelyFS reads a directory recursively and returns as model.GenericObject
func readDirRecursivelyFS(iofs fs.FS, key string) (*model.GenericObject, error) {
	data := model.GenericObject{}

	subFS, err := fs.Sub(iofs, key)
	if err != nil {
		return nil, err
	}

	err = fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}

		fileContent, err := readFileFS[model.GenericObject](subFS, path)
		if err != nil {
			return err
		}

		if len(fileContent) == 0 {
			return nil
		}

		if content, ok := fileContent[key]; ok {
			maps.Copy(data, content.(model.GenericObject))
		} else {
			ext := filepath.Ext(path)
			data[strings.TrimSuffix(path, ext)] = fileContent
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}
