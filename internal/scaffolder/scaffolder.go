package scaffolder

import (
	"os"
	"path/filepath"
	"text/template"
)

func CreateProject(name string) error {
	// Create base project folder
	if err := os.Mkdir(name, 0755); err != nil {
		return err
	}

	files := map[string]string{
		"go.mod":  `module {{.Name}}\n\ngo 1.20`,
		"main.go": `package main\n\nfunc main() {\n\tprintln("Hello from {{.Name}}")\n}`,
	}

	for fileName, content := range files {
		path := filepath.Join(name, fileName)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		tmpl, _ := template.New(fileName).Parse(content)
		tmpl.Execute(f, map[string]string{"Name": name})
		f.Close()
	}

	return nil
}
