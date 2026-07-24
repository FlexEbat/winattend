package model

type Tweaks struct {
	Explorer ExplorerTweaks

	Taskbar TaskbarTweaks

	Updates UpdateTweaks

	Privacy PrivacyTweaks

	System SystemTweaks
}

type ExplorerTweaks struct {
	ShowExtensions bool

	ShowHiddenFiles bool

	ClassicContextMenu bool
}

type TaskbarTweaks struct {
	Alignment string

	ShowSearch bool

	ShowWidgets bool

	ShowChat bool
}

type UpdateTweaks struct {
	Disable bool

	DisableDrivers bool

	DisableRestart bool
}

type PrivacyTweaks struct {
	DisableTelemetry bool

	DisableAdvertisingID bool

	DisableActivityHistory bool
}

type SystemTweaks struct {
	ShowThisPC bool

	ShowControlPanel bool

	DisableLockScreen bool
}
