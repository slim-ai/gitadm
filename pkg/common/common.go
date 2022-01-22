package common

const (
	CommandNameAdd     = "add"
	CommandNameDel     = "rm"
	CommandNameUpdate  = "update"
	CommandNameDetails = "describe"
	CommandNameList    = "ls"
)

const (
	EnvVarAPIToken         = "GITLAB_API_TOKEN"
	EnvVarConfigFilePath   = "GITLAB_CONFIG_FILE"
	EnvVarConfigHomeDir    = "XDG_CONFIG_HOME" // https://wiki.archlinux.org/index.php/XDG_Base_Directory
	EnvVarPlatformEndpoint = "GITLAB_ENDPOINT"
	EnvVarUserHomeDir      = "HOME"
)

const (
	// Default values
	DefaultToken = "YOUR_TOKEN_GOES_HERE"
)

const (
	ParamToken  = "token"
	ParamConfig = "config-file"
)
