package main

import (
	"HashTable/components/CustomIconButton"
	"HashTable/components/HashTable"
	"HashTable/components/NameList"
	"fmt"
	"image/color"
	"strings"

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

	titleLabel := canvas.NewText("PROCURAR CONTATO", color.White)
	titleLabel.TextSize = 25
	titleLabel.TextStyle.Bold = true
	titleLabel.Move(fyne.NewPos(260, 20))

	searchLabel := canvas.NewText("Buscar", color.White)
	searchLabel.TextSize = 20
	searchLabel.Move(fyne.NewPos(200, 210))

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Digite o nome para a busca...")
	searchEntry.Resize(fyne.NewSize(400, 50))
	searchEntry.Move(fyne.NewPos(200, 250))

	searchButton := widget.NewButton("Buscar", func() {
		text := searchEntry.Text
		screen1 := makeDataScreen(w, hash_table, text)
		w.SetContent(screen1)
	})
	searchButton.Importance = widget.HighImportance
	searchButton.Move(fyne.NewPos(300, 450))
	searchButton.Resize(fyne.NewSize(200, 40))

	return container.NewWithoutLayout(
		backGround,
		prevButton,
		titleLabel,
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

/***********************************TELAS INTERMEDIARIAS*************************************/

func makeDataScreen(w fyne.Window, hash_table *HashTable.Hash, nome string) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := makeSearchScreen(w, hash_table)
		w.SetContent(screen1)
	})
	prevButton.Resize(fyne.NewSize(40, 40))
	prevButton.Move(fyne.NewPos(10, 10))

	titleLabel := canvas.NewText("DADOS CADASTRADOS COM O ESSE NOME", color.White)
	titleLabel.TextSize = 25
	titleLabel.TextStyle.Bold = true
	titleLabel.Move(fyne.NewPos(150, 20))

	data, err := HashTable.BuscaHash(hash_table, nome)
	if err != nil {
		data := canvas.NewText("Não há dados cadastrados com esse nome", color.White)
		data.TextSize = 20
		data.Move(fyne.NewPos(210, 100))
		return container.NewWithoutLayout(
			backGround,
			prevButton,
			titleLabel,
			data,
		)
	}

	// Create a circle and initial label for the user
	circle := canvas.NewCircle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	circle.Resize(fyne.NewSize(100, 100))
	circle.Move(fyne.NewPos(350, 100))

	initial := canvas.NewText(string(data[0][0]), color.White)
	initial.TextSize = 50
	initial.Move(fyne.NewPos(385, 115))

	// Create a container to hold the circle and initial
	userIcon := container.NewWithoutLayout(circle, initial)

	// Create a VBox to hold the user info blocks
	vbox := container.NewVBox()

	// Iterate over the data and add user info blocks to the VBox
	for i, str := range data {
		// Split the string into name, phone, and address
		parts := strings.Split(str, ",")

		// Add the repetition indicator to the name
		if i > 0 {
			parts[0] = fmt.Sprintf("%s (%d)", parts[0], i+1)
		}

		// Create labels for the user's name, phone, and address
		nameLabel := canvas.NewText(parts[0], color.White)
		nameLabel.TextSize = 23

		phoneLabel := canvas.NewText(parts[1], color.White)
		phoneLabel.TextSize = 23

		addressLabel := canvas.NewText(parts[2], color.White)
		addressLabel.TextSize = 23

		// Create a separator line
		var separator *canvas.Line
		if i > 0 {
			separator = canvas.NewLine(color.RGBA{R: 2, G: 30, B: 80, A: 255})
			separator.StrokeWidth = 2
		} else {
			separator = canvas.NewLine(color.Transparent)
			separator.StrokeWidth = 2
		}

		space := canvas.NewRectangle(color.Transparent)
		space.Resize(fyne.NewSize(10, 5))

		// Add the user info block to the VBox
		vbox.Add(container.NewVBox(separator, space, space, space, space, container.NewCenter(nameLabel), space, container.NewCenter(phoneLabel), space, container.NewCenter(addressLabel), space, space, space, space))
	}
	// Create a scroll container for the VBox
	scroll := container.NewVScroll(vbox)
	scroll.Resize(fyne.NewSize(500, 350))
	scroll.Move(fyne.NewPos(150, 220))

	return container.NewWithoutLayout(
		backGround,
		prevButton,
		userIcon,
		titleLabel,
		scroll,
	)
}
