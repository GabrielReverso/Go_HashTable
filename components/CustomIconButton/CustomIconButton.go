package CustomIconButton

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type IconButton struct {
	widget.BaseWidget
	Icon     fyne.Resource
	OnTapped func()
}

func NewIconButton(icon fyne.Resource, tapped func()) *IconButton {
	b := &IconButton{
		Icon:     icon,
		OnTapped: tapped,
	}
	b.ExtendBaseWidget(b)
	return b
}

func (b *IconButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped()
	}
}

func (b *IconButton) CreateRenderer() fyne.WidgetRenderer {
	icon := canvas.NewImageFromResource(b.Icon)
	icon.FillMode = canvas.ImageFillOriginal

	return &iconButtonRenderer{
		button: b,
		icon:   icon,
		objects: []fyne.CanvasObject{
			icon,
		},
	}
}

type iconButtonRenderer struct {
	button  *IconButton
	icon    *canvas.Image
	objects []fyne.CanvasObject
}

func (r *iconButtonRenderer) MinSize() fyne.Size {
	return r.icon.MinSize()
}

func (r *iconButtonRenderer) Layout(size fyne.Size) {
	r.icon.Resize(size)
}

func (r *iconButtonRenderer) Refresh() {
	r.icon.Refresh()
}

func (r *iconButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *iconButtonRenderer) Destroy() {}
