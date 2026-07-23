type Scripts struct {

	System []Script

	DefaultUser []Script

	FirstLogon []Script

	UserOnce []Script
}

type Script struct {

	Name string

	Type ScriptType

	Content string
}

type ScriptType string

const (

	CMD ScriptType = "cmd"

	PowerShell ScriptType = "powershell"

	REG ScriptType = "registry"

	VBS ScriptType = "vbscript"
)
