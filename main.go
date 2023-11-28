package main

import (
	"HashTable/components/HashTable"
	"HashTable/components/MakeScreens"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {

	application := app.New()
	application.Settings().SetTheme(theme.DarkTheme())
	window := application.NewWindow("Lista Telefônica")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))
	window.SetFixedSize(true)

	hash := HashTable.CriaHash()
	/*
		HashTable.InserirDados(hash, "Gabriel Reverso", "Tel1", "End1")
		HashTable.InserirDados(hash, "Gabriel Reverso", "Tel2", "End2")
		HashTable.InserirDados(hash, "Gabriel Reverso", "Tel3", "End3") */

	screen1 := MakeScreens.MakeMainScreen(window, hash)

	window.SetContent(screen1)
	window.ShowAndRun()
}
