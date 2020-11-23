package lib

var DefaultToolKitConfig ToolKitConfig

func InitDefault() {
	DefaultToolKitConfig.CredentialDir = ".cred"
	DefaultToolKitConfig.CredentialFile = ".cred"
}
