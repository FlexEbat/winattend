package model

type Scripts struct {
	WindowsPE []Script

	Specialize []Script

	FirstLogon []Script

	DefaultUser []Script

	UserOnce []Script
}

type Script struct {
	Name string

	Type ScriptType

	Content string

	Order int
}
