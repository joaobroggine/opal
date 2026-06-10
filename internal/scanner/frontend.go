package scanner

import (
	"os"
	"path/filepath"
)

func IsFrontend(root string) (bool, error) {

	files := []string{
		"package.json",
		"next.config.js",
		"next.config.ts",
		"vite.config.js",
		"vite.config.ts",
		"webpack.config.js",
		"webpack.config.ts",
		"angular.json",
		"vue.config.js",
		"vue.config.ts",
		"svelte.config.js",
		"svelte.config.ts",
		"rollup.config.js",
		"rollup.config.ts",
		"gulpfile.js",
		"gulpfile.ts",
		"gruntfile.js",
		"gruntfile.ts",
		"ember-cli-build.js",
		"ember-cli-build.ts",
		"parcel.config.js",
		"parcel.config.ts",
		"snowpack.config.js",
		"snowpack.config.ts",
		"react-app-env.d.ts",
		"tsconfig.json",
		"jsconfig.json",
		"babel.config.js",
		"babel.config.json",
		".babelrc",
		".babelrc.js",
		".babelrc.json",
		".eslintrc",
		".eslintrc.js",
		".eslintrc.json",
		".eslintignore",
		".stylelintrc",
		".stylelintrc.js",
		".stylelintrc.json",
		".stylelintignore",
		".postcssrc",
		".postcssrc.js",
		".postcssrc.json",
		"postcss.config.js",
		"postcss.config.json",
		"tailwind.config.js",
		"tailwind.config.ts",
		"tailwind.config.cjs",
		"tailwind.config.mjs",
		"tailwind.config.json",
	}

	for _, file := range files {

		_, err := os.Stat(filepath.Join(root, file))

		if err == nil {
			return true, nil
		}
	}
	return false, nil
}

func HasNodeModules(root string) (bool, error) {

	_, err := os.Stat(filepath.Join(root, "node_modules"))
	return err == nil, nil

}