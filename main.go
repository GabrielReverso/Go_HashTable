package main

import (
	"HashTable/components/HashTable"
	"HashTable/components/MakeScreens"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {

	// Cria o aplicativo
	application := app.New()
	application.Settings().SetTheme(theme.DarkTheme())

	// Cria a janela do aplicativo
	window := application.NewWindow("Lista Hash")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))
	window.SetFixedSize(true)

	// Cria a hash
	hash := HashTable.CriaHash()

	// Obtem o conteúdo da tela inicial
	screen1 := MakeScreens.MakeMainScreen(window, hash)

	// Define o conteúdo inicial
	window.SetContent(screen1)

	// Executa o aplicativo
	window.ShowAndRun()
}
