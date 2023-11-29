package CustomIconButton

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Define a estrutura IconButton
type IconButton struct {
	widget.BaseWidget
	Icon     fyne.Resource
	OnTapped func()
}

// Função para criar um novo IconButton
func NewIconButton(icon fyne.Resource, tapped func()) *IconButton {
	b := &IconButton{
		Icon:     icon,
		OnTapped: tapped,
	}
	b.ExtendBaseWidget(b) // Estende o widget base
	return b
}

// Função para lidar com o evento de toque
func (b *IconButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped() // Chama a função OnTapped se não for nula
	}
}

// Função para criar um renderizador para o IconButton
func (b *IconButton) CreateRenderer() fyne.WidgetRenderer {
	icon := canvas.NewImageFromResource(b.Icon) // Cria uma nova imagem a partir do recurso do ícone
	icon.FillMode = canvas.ImageFillOriginal    // Define o modo de preenchimento da imagem

	// Retorna um novo iconButtonRenderer
	return &iconButtonRenderer{
		button: b,
		icon:   icon,
		objects: []fyne.CanvasObject{
			icon,
		},
	}
}

// Define a estrutura iconButtonRenderer
type iconButtonRenderer struct {
	button  *IconButton
	icon    *canvas.Image
	objects []fyne.CanvasObject
}

// Função para obter o tamanho mínimo do ícone
func (r *iconButtonRenderer) MinSize() fyne.Size {
	return r.icon.MinSize()
}

// Função para organizar o layout do ícone
func (r *iconButtonRenderer) Layout(size fyne.Size) {
	r.icon.Resize(size) // Redimensiona o ícone para o tamanho dado
}

// Função para atualizar o ícone
func (r *iconButtonRenderer) Refresh() {
	r.icon.Refresh() // Atualiza o ícone
}

// Função para obter os objetos do iconButtonRenderer
func (r *iconButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects // Retorna os objetos
}

// Função para destruir o iconButtonRenderer
func (r *iconButtonRenderer) Destroy() {}
