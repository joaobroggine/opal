package scanner

type Project struct {
	Path 	string
	Languages []string
	Frontend bool
	NodeModules bool
}

func ScanProject(root string) (*Project, error) {

	languages, err := DetectLanguages(root)
	if err != nil {
		return nil, err
	}

	var langs []string
	for lang := range languages {
		langs = append(langs, lang)
	}

	isFrontend, err := IsFrontend(root)
	if err != nil {
		return nil, err
	}

	hasModules, err := HasNodeModules(root)
	if err != nil {
		return nil, err
	}

	project := &Project{
		Path:         root,
		Languages:    langs,
		Frontend:     isFrontend,
		NodeModules:  hasModules,
	}

	return project, nil
}