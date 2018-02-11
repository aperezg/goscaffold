package goscaffold

import (
	"flag"
	"fmt"
	"github.com/aperezg/goscaffold/data"
	"os"
)

var Client consoleClient

type consoleClient struct{}

func (c *consoleClient) Stdin() (setting *Settings) {
	c.configureHelp()

	command := flag.NewFlagSet("create-project", flag.ExitOnError)

	enableGitlabCI := command.Bool("gitlabci", false, "Create a .gitlab-ci with a valid pipeline")
	enableDocker := command.Bool("docker", false, "Create a project with use a docker container for go project")
	workspace := command.String("workspace", "$GOPATH/src", "Your golang workspace")

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	command.Parse(os.Args[4:])
	if command.Parsed() {
		applicationName, importPath := c.extractImportPath(workspace)
		setting = NewSettings(
			applicationName,
			importPath,
			*enableGitlabCI,
			*enableDocker,
		)
	}

	return
}

func (c *consoleClient) extractImportPath(workspace *string) (string, string) {
	args := os.Args[2:4]
	var namespace, applicationName = args[0], args[1]
	importPath := os.ExpandEnv(*workspace) + string(os.PathSeparator) + namespace + string(os.PathSeparator) + applicationName
	return applicationName, importPath
}

func (c *consoleClient) configureHelp() {
	flag.Usage = func() {
		d, err := data.Asset("doc/help")
		if err != nil {
			fmt.Println("Can't load the help, go to de 'https://github.com/aperezg/goscaffold' for more information")
			os.Exit(1)
		}

		fmt.Println(string(d))
	}
}
