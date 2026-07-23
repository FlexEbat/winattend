package model

type Windows struct {
	Edition      Edition
	Architecture Architecture

	ComputerName string
	TimeZone     string

	ProductKey Secret
}
