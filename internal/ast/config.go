package ast

type Config struct {
	Metadata Metadata

	Windows Windows

	Language Language

	Disk Disk

	Users []User

	Network Network

	Security Security

	Tweaks Tweaks

	Apps Apps

	Scripts Scripts
}
