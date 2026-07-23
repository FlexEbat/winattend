package ast

type Tweaks struct {

	Explorer ExplorerTweaks

	Update UpdateTweaks

	Taskbar TaskbarTweaks

	Privacy PrivacyTweaks

	System SystemTweaks
}

type ExplorerTweaks struct {

	ShowExtensions bool

	ShowHidden bool

	ClassicContextMenu bool
}

type UpdateTweaks struct {

	Disable bool

	PreventRestart bool
}
