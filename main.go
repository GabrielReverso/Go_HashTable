package main

import (
	"HashTable/components/CustomIconButton"
	"HashTable/components/HashTable"
	"HashTable/components/NameList"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	application := app.New()
	application.Settings().SetTheme(theme.DarkTheme())
	window := application.NewWindow("Lista Telefônica")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))
	window.SetFixedSize(true)

	hash := HashTable.CriaHash()

	screen1 := makeMainScreen(window, hash)

	window.SetContent(screen1)
	window.ShowAndRun()
}

func makeMainScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	title := canvas.NewText("LISTA TELEFÔNICA", color.White)
	title.TextSize = 26
	title.TextStyle.Bold = true
	title.Move(fyne.NewPos(255, 10))

	icon := canvas.NewImageFromResource(theme.StorageIcon())
	icon.FillMode = canvas.ImageFillOriginal
	icon.Resize(fyne.NewSize(26, 26))
	icon.Move(fyne.NewPos(505, 17))

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	addButton := CustomIconButton.NewIconButton(theme.ContentAddIcon(), func() {
		screen2 := makeAddScreen(w, hash_table)
		w.SetContent(screen2)
	})
	addButton.Resize(fyne.NewSize(40, 40))
	addButton.Move(fyne.NewPos(375, 550))

	removeButton := CustomIconButton.NewIconButton(theme.DeleteIcon(), func() {
		screen2 := makeRemoveScreen(w, hash_table)
		w.SetContent(screen2)
	})
	removeButton.Resize(fyne.NewSize(40, 40))
	removeButton.Move(fyne.NewPos(100, 550))

	searchButton := CustomIconButton.NewIconButton(theme.SearchIcon(), func() {
		screen2 := makeSearchScreen(w, hash_table)
		w.SetContent(screen2)
	})
	searchButton.Resize(fyne.NewSize(40, 40))
	searchButton.Move(fyne.NewPos(650, 550))

	names, err := HashTable.BuscaTodosHash(hash_table)
	if err != nil {
		data := canvas.NewText("Não há dados cadastrados", color.White)
		data.TextSize = 20
		data.Move(fyne.NewPos(265, 100))

		return container.NewWithoutLayout(
			backGround,
			title,
			icon,
			data,
			addButton,
			removeButton,
			searchButton,
		)
	} else {
		data := NameList.CreateButtons(names)
		maxScroll := container.NewGridWrap(fyne.NewSize(790, 490))
		maxScroll.Add(container.NewVScroll(data))
		maxScroll.Move(fyne.NewPos(0, 60))

		return container.NewWithoutLayout(
			backGround,
			title,
			icon,
			maxScroll,
			addButton,
			removeButton,
			searchButton,
		)
	}

}

func makeAddScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := makeMainScreen(w, hash_table)
		w.SetContent(screen1)
	})
	prevButton.Resize(fyne.NewSize(40, 40))
	prevButton.Move(fyne.NewPos(10, 10))

	titleLabel := canvas.NewText("ADICIONAR NOVO CONTATO", color.White)
	titleLabel.TextSize = 25
	titleLabel.TextStyle.Bold = true
	titleLabel.Move(fyne.NewPos(220, 20))

	nameLabel := canvas.NewText("Nome", color.White)
	nameLabel.TextSize = 20
	nameLabel.Move(fyne.NewPos(200, 100))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Digite o nome...")
	nameEntry.Resize(fyne.NewSize(400, 50))
	nameEntry.Move(fyne.NewPos(200, 130))

	phoneLabel := canvas.NewText("Telefone", color.White)
	phoneLabel.TextSize = 20
	phoneLabel.Move(fyne.NewPos(200, 200))

	phoneEntry := widget.NewEntry()
	phoneEntry.SetPlaceHolder("Digite o telefone...")
	phoneEntry.Resize(fyne.NewSize(400, 50))
	phoneEntry.Move(fyne.NewPos(200, 230))

	addressLabel := canvas.NewText("Endereço", color.White)
	addressLabel.TextSize = 20
	addressLabel.Move(fyne.NewPos(200, 300))

	addressEntry := widget.NewEntry()
	addressEntry.SetPlaceHolder("Digite o endereço...")
	addressEntry.Resize(fyne.NewSize(400, 50))
	addressEntry.Move(fyne.NewPos(200, 330))

	confirmButton := widget.NewButton("Adicionar", func() {
		name := nameEntry.Text
		phone := phoneEntry.Text
		address := addressEntry.Text

		if name == "" || phone == "" || address == "" {
			dialog.ShowInformation("Alerta!", "Preencha todos os campos para prosseguir!", w)
		} else {
			HashTable.InserirDados(hash_table, name, phone, address)
			content := container.NewVBox(
				widget.NewLabel("Sucesso! Dados inseridos na lista."),
				widget.NewLabel("Deseja inserir mais dados?\n"),
			)

			dialog.ShowCustomConfirm("Confirmação", "Sim", "Não", content, func(response bool) {
				if !response {
					screen1 := makeMainScreen(w, hash_table)
					w.SetContent(screen1)
				} else {
					screen1 := makeAddScreen(w, hash_table)
					w.SetContent(screen1)
				}
			}, w)
		}
	})
	confirmButton.Importance = widget.HighImportance
	confirmButton.Move(fyne.NewPos(300, 450))
	confirmButton.Resize(fyne.NewSize(200, 40))

	return container.NewWithoutLayout(
		backGround,
		prevButton,
		titleLabel,
		nameLabel,
		nameEntry,
		phoneLabel,
		phoneEntry,
		addressLabel,
		addressEntry,
		confirmButton,
	)
}

func makeSearchScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := makeMainScreen(w, hash_table)
		w.SetContent(screen1)
	})
	prevButton.Resize(fyne.NewSize(40, 40))
	prevButton.Move(fyne.NewPos(10, 10))

	searchLabel := canvas.NewText("Buscar", color.White)
	searchLabel.TextSize = 20
	searchLabel.Move(fyne.NewPos(200, 210))

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Digite o nome para a busca...")
	searchEntry.Resize(fyne.NewSize(400, 50))
	searchEntry.Move(fyne.NewPos(200, 250))

	searchButton := widget.NewButton("Buscar", func() {
		// Implement search logic here
	})
	searchButton.Importance = widget.HighImportance
	searchButton.Move(fyne.NewPos(300, 450))
	searchButton.Resize(fyne.NewSize(200, 40))

	return container.NewWithoutLayout(
		backGround,
		prevButton,
		searchLabel,
		searchEntry,
		searchButton,
	)
}

func makeRemoveScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := makeMainScreen(w, hash_table)
		w.SetContent(screen1)
	})
	prevButton.Resize(fyne.NewSize(50, 50))
	prevButton.Move(fyne.NewPos(100, 100))

	return container.NewWithoutLayout(
		backGround,
		prevButton,
	)
}
