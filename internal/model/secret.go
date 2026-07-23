package model

// Secret описывает секрет без привязки к конкретному backend.
//
// Примеры:
//
//	password: "123456"
//	password: env(ADMIN_PASSWORD)
//	password: sops(admin.password)
//	password: vault(secret/windows/admin)
type Secret struct {
	Provider SecretProvider

	// Reference используется для env/sops/vault/gpg...
	Reference string

	// Value используется только для SecretPlain.
	Value string
}

func (s Secret) IsEmpty() bool {
	return s.Provider == "" &&
		s.Reference == "" &&
		s.Value == ""
}

func (s Secret) IsPlain() bool {
	return s.Provider == SecretPlain
}
