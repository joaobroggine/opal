package scanner

import (
	"os"
	"path/filepath"
)

var Extensions = map[string]string{
	".go":   "Go",
	".py":   "Python",
	".js":   "JavaScript",
	".jsx":  "React JS",
	".ts":   "TypeScript",
	".tsx":  "React TS",
	".java": "Java",
	".rb":   "Ruby",
	".r":    "R",
	".php":  "PHP",
	".phtml":"PHP",
	".html": "HTML",
	".css":  "CSS",
	".cs":   "C#",
	".cpp":  "C++",
	".c":    "C",
	".rs":   "Rust",
	".swift":"Swift",
	".kt":   "Kotlin",
	".sql":  "SQL",
	".cbl": "COBOL",
	".cob": "COBOL",
	".cpy": "COBOL",
	".asm":  "Assembly",
	".pl":   "Perl",
	".lua":  "Lua",
	".scala":"Scala",
	".dart": "Dart",
}

func DetectLanguages(root string) (map[string]bool, error) {

	found := make(map[string]bool)

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

		ext := filepath.Ext(path)

		if lang, ok := Extensions[ext]; ok {
			found[lang] = true
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return found, nil
}