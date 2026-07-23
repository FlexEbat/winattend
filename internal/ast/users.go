package ast

type User struct {

	Name string

	DisplayName string

	Password Secret

	Admin bool

	AutoLogon bool
}
