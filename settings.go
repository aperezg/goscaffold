package goscaffold

type Settings struct {
	ApplicationName string
	ImportPath      string
	Namespace       string
	EnableGitlabCI  bool
	EnableDocker    bool
}

func NewSettings(
	applicationName string,
	importPath string,
	namespace string,
	enableGitlabCI bool,
	enableDocker bool,
) *Settings {

	return &Settings{
		ApplicationName: applicationName,
		ImportPath:      importPath,
		Namespace:       namespace,
		EnableGitlabCI:  enableGitlabCI,
		EnableDocker:    enableDocker,
	}
}
