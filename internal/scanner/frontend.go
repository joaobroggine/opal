package scanner

import (
	"os"
	"path/filepath"
)

func IsFrontend(root string) (bool, error) {

	files := map[string]bool{
		"package.json": true,
		"next.config.js": true,
		"next.config.ts": true,
		"vite.config.js": true,
		"vite.config.ts": true,
		"webpack.config.js": true,
		"webpack.config.ts": true,
		"angular.json": true,
		"vue.config.js": true,
		"vue.config.ts": true,
		"svelte.config.js": true,
		"svelte.config.ts": true,
		"rollup.config.js": true,
		"rollup.config.ts": true,
		"gulpfile.js": true,
		"gulpfile.ts": true,
		"gruntfile.js": true,
		"gruntfile.ts": true,
		"ember-cli-build.js": true,
		"ember-cli-build.ts": true,
		"parcel.config.js": true,
		"parcel.config.ts": true,
		"snowpack.config.js": true,
		"snowpack.config.ts": true,
		"react-app-env.d.ts": true,
		"tsconfig.json": true,
		"jsconfig.json": true,
		"babel.config.js": true,
		"babel.config.json": true,
		".babelrc": true,
		".babelrc.js": true,
		".babelrc.json": true,
		".eslintrc": true,
		".eslintrc.js": true,
		".eslintrc.json": true,
		".eslintignore": true,
		".stylelintrc": true,
		".stylelintrc.js": true,
		".stylelintrc.json": true,
		".stylelintignore": true,
		".postcssrc": true,
		".postcssrc.js": true,
		".postcssrc.json": true,
		"postcss.config.js": true,
		"postcss.config.json": true,
		"tailwind.config.js": true,
		"tailwind.config.ts": true,
		"tailwind.config.cjs": true,
		"tailwind.config.mjs": true,
		"tailwind.config.json": true,
	}

	found := false

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

		if files[info.Name()] {
			found = true
			return filepath.SkipAll
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return found, nil

}

func HasNodeModules(root string) (bool, error) {

	found := false

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "node_modules" {
			found = true
			return filepath.SkipAll
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return found, nil

}