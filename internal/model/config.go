package model

type Config struct {
	Windows  Windows
	Language Language
	Disk     Disk

	Users []User

	Network Network

	Security Security

	Tweaks Tweaks

	Apps Apps

	Scripts Scripts
}
