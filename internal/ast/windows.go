package ast

type Windows struct {
	Edition Edition

	Architecture Architecture

	ProductKey Secret

	ComputerName string

	TimeZone string
}
