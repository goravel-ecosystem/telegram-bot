package support

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadDocs(repoPath string) (map[string]string, error) {
	docs := make(map[string]string)

	// Walk through the directory
	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}
		if strings.HasSuffix(info.Name(), ".md") {
			// Read file contents
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			path, err = filepath.Rel(repoPath, path)
			if err != nil {
				return err
			}

			// Store in map with file path as key
			docs[path] = string(content)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return docs, nil
}
