package ast

type Network struct {

	Wifi []WifiNetwork
}

type WifiNetwork struct {

	SSID string

	Password Secret

	Hidden bool

	Authentication WifiAuthentication
}

type WifiAuthentication string

const (

	Open WifiAuthentication = "open"

	WPA2 WifiAuthentication = "wpa2"

	WPA3 WifiAuthentication = "wpa3"
)
