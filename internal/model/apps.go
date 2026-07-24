package model

type Apps struct {
	Install []Package

	Remove []Package
}

type Package struct {
	ID string

	Source string

	Version string
}
