package model

type Edition string

const (
	EditionHome       Edition = "home"
	EditionPro        Edition = "pro"
	EditionEnterprise Edition = "enterprise"
	EditionEducation  Edition = "education"
)

type Architecture string

const (
	ArchitectureAMD64 Architecture = "amd64"
	ArchitectureARM64 Architecture = "arm64"
	ArchitectureX86   Architecture = "x86"
)

type DiskLayout string

const (
	DiskLayoutAuto DiskLayout = "auto"
	DiskLayoutGPT  DiskLayout = "gpt"
	DiskLayoutMBR  DiskLayout = "mbr"
)

type WifiAuthentication string

const (
	WifiOpen WifiAuthentication = "open"
	WifiWPA2 WifiAuthentication = "wpa2"
	WifiWPA3 WifiAuthentication = "wpa3"
)

type ScriptType string

const (
	ScriptCMD        ScriptType = "cmd"
	ScriptPowerShell ScriptType = "powershell"
	ScriptRegistry   ScriptType = "registry"
	ScriptVBScript   ScriptType = "vbscript"
)

type SecretProvider string

const (
	SecretPlain SecretProvider = "plain"
	SecretEnv   SecretProvider = "env"
	SecretSOPS  SecretProvider = "sops"
	SecretVault SecretProvider = "vault"
	SecretGPG   SecretProvider = "gpg"
)
