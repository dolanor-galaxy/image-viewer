package main

import (
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (a *App) loadSettingsUI() {
	winSettings := a.app.NewWindow("Settings")

	themeSelector := widget.NewSelect([]string{"Light", "Dark"}, func(selected string) {
		switch selected {
		case "Light":
			a.app.Settings().SetTheme(theme.LightTheme())
		case "Dark":
			a.app.Settings().SetTheme(theme.DarkTheme())
		}
		a.app.Preferences().SetString("Theme", selected)
	})
	themeSelector.SetSelected(a.app.Preferences().StringWithFallback("Theme", "Dark"))

	winSettings.SetContent(widget.NewVBox(
		widget.NewHBox(
			widget.NewLabel("Theme"),
			themeSelector,
		),
		layout.NewSpacer(),
		widget.NewLabel("v1.0 | License: MIT"),
		widget.NewHyperlink("Github (source code and more information)", parseURL("https://github.com/Palexer/image-viewer")),
	))
	winSettings.Show()
}

func (a *App) loadPreferences() {
	// theme
	switch a.app.Preferences().StringWithFallback("Theme", "Dark") {
	case "Light":
		a.app.Settings().SetTheme(theme.LightTheme())
	case "Dark":
		a.app.Settings().SetTheme(theme.DarkTheme())
	}

	// show/hide statusbar
	if a.app.Preferences().BoolWithFallback("statusBarVisible", true) == false {
		a.statusBar.Hide()
	}
}
