package model

type Language struct {
	Display string
	System  string
	Input   string
	Locale  string
	Geo     string

	Fallback []string
}
