package model

type Network struct {
	Wifi []WifiNetwork
}

type WifiNetwork struct {
	SSID string

	Password Secret

	Hidden bool

	Authentication WifiAuthentication
}
