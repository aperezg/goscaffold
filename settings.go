package goscaffold

type Settings struct {
	ApplicationName string
	ImportPath      string
	EnableGitlabCI  bool
	EnableDocker    bool
}

func NewSettings(
	applicationName string,
	importPath string,
	enableGitlabCI bool,
	enableDocker bool,
) *Settings {

	return &Settings{
		ApplicationName: applicationName,
		ImportPath:      importPath,
		EnableGitlabCI:  enableGitlabCI,
		EnableDocker:    enableDocker,
	}
}
