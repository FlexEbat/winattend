package model

type Security struct {
	UAC bool

	Defender bool

	SmartScreen bool

	BitLocker bool

	TPMBypass bool

	SecureBootBypass bool
}
