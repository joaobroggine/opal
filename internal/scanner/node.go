package scanner

import (
	"os"
	"path/filepath"
)

func FindNodeProjects(root string) ([]string, error) {

	paths := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			switch info.Name() {
			case "node_modules", ".git", ".next", "dist", "build":
				return filepath.SkipDir
			}
			return nil
		}

		if info.Name() == "package.json" {
			dir := filepath.Dir(path)
			paths = append(paths, dir)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}