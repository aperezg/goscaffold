package goscaffold

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Generator struct {
	settings  *Settings
	templates []templateStructure
}

type templateStructure struct {
	Path     string `yaml:"path"`
	FileName string `yaml:"fileName"`
}

const TEMPLATES_DIRECTORY = "templates"
const STRUCTURE_FILE_NAME = "structure.yml"

func NewGenerator(settings *Settings) *Generator {
	g := &Generator{settings: settings}
	g.getTemplateStructure()
	return g
}

func (g *Generator) getTemplateStructure() {
	structureFile := TEMPLATES_DIRECTORY + string(os.PathSeparator) + STRUCTURE_FILE_NAME
	yamlFile, err := ioutil.ReadFile(structureFile)
	if err != nil {
		log.Fatalf("Error reading the structure file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &g.templates)
	if err != nil {
		log.Fatalf("Error parsing the structure file: %v", err)
	}
}

func (g *Generator) Scaffold() error {
	fmt.Println(fmt.Sprintf("Generating the directory structure %s ...", g.settings.ImportPath))
	g.createDirectory(g.settings.ImportPath)
	fmt.Println("Generating files project...")
	g.generateFiles()
	return nil
}

func (g *Generator) createDirectory(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatalf("Impossible create the directory {%s}: %v", g.settings.ImportPath, err)
	}
}

func (g *Generator) generateFiles() {
	for _, t := range g.templates {

		templateFilePath := TEMPLATES_DIRECTORY + string(os.PathSeparator) + t.FileName
		finalFileName := strings.TrimRight(strings.TrimLeft(t.FileName, "_"), ".tmpl")
		outputPath := g.settings.ImportPath + string(os.PathSeparator)
		if t.Path != "" {
			outputPath += t.Path + string(os.PathSeparator)
			g.createDirectory(outputPath)
		}
		outputFile := outputPath + finalFileName
		fmt.Println(fmt.Sprintf("Generating file %s ...", outputFile))

		b, err := ioutil.ReadFile(templateFilePath)
		if err != nil {
			log.Fatalf("Can't open %s: %v", templateFilePath, err)
		}

		tmpl := template.New(finalFileName)
		tmpl, err = tmpl.Parse(string(b))
		if err != nil {
			log.Fatalf("Error parsing the file %s: %v", templateFilePath, err)
		}

		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("Error creating the file %s: %v", outputFile, err)
		}

		if err := tmpl.Execute(f, g.settings); err != nil {
			log.Fatalf("Error saving the new file %s: %v", outputFile, err)
		}

		f.Close()
	}
}
