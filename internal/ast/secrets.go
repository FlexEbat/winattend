package ast

type Secret struct {

	Provider SecretProvider

	Value string
}

type SecretProvider string

const (

	Plain SecretProvider = "plain"

	Env SecretProvider = "env"

	SOPS SecretProvider = "sops"

	Vault SecretProvider = "vault"

	GPG SecretProvider = "gpg"
)
