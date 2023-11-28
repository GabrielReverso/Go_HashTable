package MakeScreens

import (
	"HashTable/components/CustomIconButton"
	"HashTable/components/HashTable"
	"math/rand"
	"sort"
	"time"

	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeMainScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

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
		screen2 := MakeAddScreen(w, hash_table)
		w.SetContent(screen2)
	})
	addButton.Resize(fyne.NewSize(40, 40))
	addButton.Move(fyne.NewPos(375, 550))

	removeButton := CustomIconButton.NewIconButton(theme.DeleteIcon(), func() {
		screen2 := MakeRemoveScreen(w, hash_table)
		w.SetContent(screen2)
	})
	removeButton.Resize(fyne.NewSize(40, 40))
	removeButton.Move(fyne.NewPos(100, 550))

	searchButton := CustomIconButton.NewIconButton(theme.SearchIcon(), func() {
		screen2 := MakeSearchScreen(w, hash_table)
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
		data := CreateButtons(w, hash_table, names)
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

func MakeAddScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := MakeMainScreen(w, hash_table)
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
					screen1 := MakeMainScreen(w, hash_table)
					w.SetContent(screen1)
				} else {
					screen1 := MakeAddScreen(w, hash_table)
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

func MakeSearchScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := MakeMainScreen(w, hash_table)
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
		if text == "" {
			dialog.ShowInformation("Alerta!", "Preencha todos os campos para prosseguir!", w)
		} else {
			screen1 := MakeDataScreen(w, hash_table, text, "", "", false)
			w.SetContent(screen1)
		}
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

func MakeRemoveScreen(w fyne.Window, hash_table *HashTable.Hash) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := MakeMainScreen(w, hash_table)
		w.SetContent(screen1)
	})
	prevButton.Resize(fyne.NewSize(40, 40))
	prevButton.Move(fyne.NewPos(10, 10))

	titleLabel := canvas.NewText("DIGITE O NOME DO CONTATO", color.White)
	titleLabel.TextSize = 25
	titleLabel.TextStyle.Bold = true
	titleLabel.Move(fyne.NewPos(233, 20))

	searchLabel := canvas.NewText("Qual nome deseja remover?", color.White)
	searchLabel.TextSize = 20
	searchLabel.Move(fyne.NewPos(200, 210))

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Digite o nome para procurar...")
	searchEntry.Resize(fyne.NewSize(400, 50))
	searchEntry.Move(fyne.NewPos(200, 250))

	searchButton := widget.NewButton("Buscar", func() {
		text := searchEntry.Text
		if text == "" {
			dialog.ShowInformation("Alerta!", "Preencha todos os campos para prosseguir!", w)
		} else {
			screen1 := MakeSelectScreen(w, hash_table, text)
			w.SetContent(screen1)
		}
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

/***********************************TELAS INTERMEDIARIAS*************************************/

func MakeDataScreen(w fyne.Window, hash_table *HashTable.Hash, nome string, telefone string, endereco string, especifico bool) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		if especifico {
			screen1 := MakeMainScreen(w, hash_table)
			w.SetContent(screen1)
		} else {
			screen1 := MakeSearchScreen(w, hash_table)
			w.SetContent(screen1)
		}
	})
	prevButton.Resize(fyne.NewSize(40, 40))
	prevButton.Move(fyne.NewPos(10, 10))

	titleLabel := canvas.NewText("DADOS CADASTRADOS COM O ESSE NOME", color.White)
	titleLabel.TextSize = 25
	titleLabel.TextStyle.Bold = true
	titleLabel.Move(fyne.NewPos(150, 20))

	var data []string
	var err error
	if especifico {
		data, err = HashTable.BuscaEspecificoHash(hash_table, nome, telefone, endereco)
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
	} else {
		data, err = HashTable.BuscaHash(hash_table, nome)
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
	}

	circle := canvas.NewCircle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	circle.Resize(fyne.NewSize(100, 100))
	circle.Move(fyne.NewPos(350, 100))

	initial := canvas.NewText(string(data[0][0]), color.White)
	initial.TextSize = 50
	initial.Move(fyne.NewPos(382, 115))

	userIcon := container.NewWithoutLayout(circle, initial)

	vbox := container.NewVBox()

	for i, str := range data {
		parts := strings.Split(str, ",")
		if i > 0 {
			parts[0] = fmt.Sprintf("%s (%d)", parts[0], i+1)
		}
		nameLabel := canvas.NewText(parts[0], color.White)
		nameLabel.TextSize = 23

		phoneLabel := canvas.NewText(parts[1], color.White)
		phoneLabel.TextSize = 23

		addressLabel := canvas.NewText(parts[2], color.White)
		addressLabel.TextSize = 23

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
		vbox.Add(container.NewVBox(separator, space, space, space, space, container.NewCenter(nameLabel), space, container.NewCenter(phoneLabel), space, container.NewCenter(addressLabel), space, space, space, space))
	}
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

func MakeSelectScreen(w fyne.Window, hash_table *HashTable.Hash, nome string) fyne.CanvasObject {

	backGround := canvas.NewRectangle(color.RGBA{R: 2, G: 20, B: 35, A: 255})
	backGround.Resize(fyne.NewSize(805, 605))
	backGround.Move(fyne.NewPos(-5, -5))

	prevButton := CustomIconButton.NewIconButton(theme.NavigateBackIcon(), func() {
		screen1 := MakeRemoveScreen(w, hash_table)
		w.SetContent(screen1)
	})
	prevButton.Resize(fyne.NewSize(40, 40))
	prevButton.Move(fyne.NewPos(10, 10))

	titleLabel := canvas.NewText("SELECIONE PARA REMOVER CONTATO", color.White)
	titleLabel.TextSize = 25
	titleLabel.TextStyle.Bold = true
	titleLabel.Move(fyne.NewPos(185, 20))

	data, err := HashTable.BuscaHash(hash_table, nome)
	fmt.Println(data, err)
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

	circle := canvas.NewCircle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	circle.Resize(fyne.NewSize(100, 100))
	circle.Move(fyne.NewPos(350, 75))

	initial := canvas.NewText(string(data[0][0]), color.White)
	initial.TextSize = 50
	initial.Move(fyne.NewPos(381, 90))

	deleteAllButton := widget.NewButton("Deletar todos", func() {
		HashTable.DeleteAllHash(hash_table, nome)
		screen := MakeMainScreen(w, hash_table)
		w.SetContent(screen)
	})
	deleteAllButton.Importance = widget.HighImportance
	deleteAllButton.Resize(fyne.NewSize(120, 30))
	deleteAllButton.Move(fyne.NewPos(340, 530))

	userIcon := container.NewWithoutLayout(circle, initial)

	vbox := container.NewVBox()

	for i, str := range data {

		parts := strings.Split(str, ",")

		if i > 0 {
			parts[0] = fmt.Sprintf("%s (%d)", parts[0], i+1)
		}

		nameLabel := canvas.NewText(parts[0], color.White)
		nameLabel.TextSize = 23

		phoneLabel := canvas.NewText(parts[1], color.White)
		phoneLabel.TextSize = 23

		addressLabel := canvas.NewText(parts[2], color.White)
		addressLabel.TextSize = 23

		deleteButton := CustomIconButton.NewIconButton(theme.DeleteIcon(), func() {
			name := parts[0]
			if strings.HasSuffix(parts[0], ")") {
				lastIndex := strings.LastIndex(parts[0], " ")
				if lastIndex != -1 {
					name = parts[0][:lastIndex]
				}
			}
			HashTable.DeleteHash(hash_table, name, parts[1])
			screen := MakeMainScreen(w, hash_table)
			w.SetContent(screen)
		})
		deleteButton.Resize(fyne.NewSize(60, 60))

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

		vbox.Add(container.NewVBox(separator, space, space, space, space, container.NewCenter(nameLabel), space, container.NewCenter(phoneLabel), space, container.NewCenter(addressLabel), space, container.NewCenter(deleteButton), space, space, space, space))
	}

	scroll := container.NewVScroll(vbox)
	scroll.Resize(fyne.NewSize(500, 300))
	scroll.Move(fyne.NewPos(150, 195))

	return container.NewWithoutLayout(
		backGround,
		prevButton,
		userIcon,
		deleteAllButton,
		titleLabel,
		scroll,
	)
}

/**********************************AUXILIARES***************************************/

type buttonLayout struct {
}

func (l *buttonLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 40) // set the size of the button
}

func (l *buttonLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	objects[0].Resize(fyne.NewSize(40, 40))  // circle
	objects[1].Resize(objects[1].MinSize())  // label
	objects[2].Resize(fyne.NewSize(700, 40)) // button

	objects[0].Move(fyne.NewPos(15, 0))
	objects[1].Move(fyne.NewPos(26, 3))
	objects[2].Move(fyne.NewPos(75, 0))
}

func (l *buttonLayout) DesiredSize(objects []fyne.CanvasObject, size fyne.Size) fyne.Size {
	return size
}

var colors = []color.RGBA{
	{R: 0, G: 0, B: 150, A: 255},
	{R: 0, G: 0, B: 200, A: 255},
	{R: 0, G: 120, B: 200, A: 255},
	{R: 30, G: 144, B: 255, A: 255},
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomColor() color.RGBA {
	return colors[rand.Intn(len(colors))]
}

func CreateButtons(w fyne.Window, hash_table *HashTable.Hash, data []string) *fyne.Container {

	// Organize o slice em ordem alfabética
	sort.Strings(data)

	// Crie um layout para os botões
	var objects []fyne.CanvasObject
	if len(data) > 0 {
		objects = make([]fyne.CanvasObject, len(data)*2-1)
	} else {
		//objects = make([]fyne.CanvasObject, 0)
		data := canvas.NewText("Não há dados cadastrados", color.White)
		data.TextSize = 20
		data.Move(fyne.NewPos(265, 100))
		return container.NewWithoutLayout(
			data,
		)
	}

	prevName := ""
	i := 1

	for j, str := range data {

		parts := strings.Split(str, ",")

		nome := parts[0]
		telefone := parts[1]
		endereco := parts[2]

		if nome == prevName {
			i++
			parts[0] = fmt.Sprintf("%s (%d)", nome, i)
		} else {
			i = 1
		}

		prevName = nome

		button := widget.NewButton(parts[0], func() {
			screen := MakeDataScreen(w, hash_table, nome, telefone, endereco, true)
			w.SetContent(screen)
		})

		circle := canvas.NewCircle(randomColor())
		circle.StrokeWidth = 0
		circle.StrokeColor = color.Black

		label := canvas.NewText(strings.ToUpper(string(str[0])), color.White)
		label.TextSize = 24
		label.TextStyle.Bold = true

		container := fyne.NewContainerWithLayout(&buttonLayout{}, circle, label, button)
		objects[j*2] = container

		if j < len(data)-1 {
			spacer := canvas.NewRectangle(color.Transparent)
			spacer.SetMinSize(fyne.NewSize(1, 3)) // 3 pixels of height
			objects[j*2+1] = spacer
		}
	}

	return container.NewVBox(objects...)
}
