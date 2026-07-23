package model

type User struct {
	Name string

	DisplayName string

	Password Secret

	Admin bool

	AutoLogon bool

	PasswordNeverExpires bool

	Groups []string
}
