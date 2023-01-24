package main

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	CARD_WIDTH                   = 60
	CARD_HEIGHT                  = 90
	HAND_CARD_OVERLAP            = 35
	TRICK_CARD_OVERLAP           = 5
	CARD_LABEL_SPACING           = 2
	CARD_COUNTERCLOCKWISE_OFFSET = 5
)

var (
	BACK_COLOR = color.RGBA{0, 160, 160, 255}
	RED        = color.RGBA{255, 0, 0, 255}
)

type handWidget struct {
	widget.BaseWidget
	visibleCards []*card
	hiddenCards  []*card
	horizontal   bool
	cardHandler  func(card *card)
}

func NewHandWidget(visibleCards []*card, hiddenCards []*card, horizontal bool, fn func(card *card)) *handWidget {
	hand := &handWidget{}
	hand.ExtendBaseWidget(hand)
	hand.horizontal = horizontal
	hand.cardHandler = fn
	hand.SetHandwidget(visibleCards, hiddenCards)

	return hand
}

func (p *handWidget) Tapped(event *fyne.PointEvent) {
	cardsCount := len(p.visibleCards) + len(p.hiddenCards)
	handLength := (cardsCount-1)*HAND_CARD_OVERLAP + CARD_WIDTH

	if event.Position.X < 0 || event.Position.X > float32(handLength) {
		return
	}
	index := int(math.Min(math.Floor(float64(event.Position.X/HAND_CARD_OVERLAP)), float64(len(p.visibleCards)-1)))

	p.cardHandler(p.visibleCards[index])
}

func (p *handWidget) CreateRenderer() fyne.WidgetRenderer {
	return newHandWidgetRenderer(p)
}

func (p *handWidget) SetHandwidget(visibleCards []*card, hiddenCards []*card) {
	p.visibleCards = visibleCards
	p.hiddenCards = hiddenCards
	p.Refresh()
}

type handWidgetRenderer struct {
	h              *handWidget
	background     *canvas.Rectangle
	paintableCards []fyne.CanvasObject
}

func newHandWidgetRenderer(h *handWidget) *handWidgetRenderer {
	render := &handWidgetRenderer{}
	render.h = h
	render.background = canvas.NewRectangle(color.Transparent)
	render.Refresh()

	return render
}

func (p *handWidgetRenderer) Layout(s fyne.Size) {
	p.background.Resize(s)
}

func (p *handWidgetRenderer) MinSize() fyne.Size {
	w := CARD_WIDTH + 7*HAND_CARD_OVERLAP
	h := CARD_HEIGHT
	if !p.h.horizontal {
		return fyne.Size{Width: float32(h), Height: float32(w)}
	}
	return fyne.Size{Width: float32(w), Height: float32(h)}
}

func (p *handWidgetRenderer) Objects() []fyne.CanvasObject {
	result := []fyne.CanvasObject{p.background}

	return append(result, p.paintableCards...)
}

func (p *handWidgetRenderer) Destroy() {}

func (p *handWidgetRenderer) Refresh() {
	offset := 0
	p.paintableCards = []fyne.CanvasObject{}
	cardHorizontal := !p.h.horizontal
	for _, v := range p.h.hiddenCards {
		paintableCard := CreatePaintableCard(v, false, cardHorizontal)
		// paintableCard := NewCardWidget(v, false, cardHorizontal)
		paintableCard.Move(newPosition(offset, p.h.horizontal))
		offset += HAND_CARD_OVERLAP
		p.paintableCards = append(p.paintableCards, paintableCard)
	}
	for _, v := range p.h.visibleCards {
		paintableCard := CreatePaintableCard(v, true, cardHorizontal)
		// paintableCard := NewCardWidget(v, true, cardHorizontal)
		paintableCard.Move(newPosition(offset, p.h.horizontal))
		offset += HAND_CARD_OVERLAP
		p.paintableCards = append(p.paintableCards, paintableCard)
	}
	canvas.Refresh(p.h)
}

func newPosition(offset int, horizontal bool) fyne.Position {
	if horizontal {
		return fyne.NewPos(float32(offset), 0)
	}
	return fyne.NewPos(0, float32(offset))
}

func CreatePaintableCard(card *card, faceUp bool, horizontal bool) *fyne.Container {
	var rectangle *canvas.Rectangle
	var labelTop, labelBottom *canvas.Text
	if faceUp {
		rectangle = canvas.NewRectangle(color.White)
		labelTop = canvas.NewText(card.String(), getColorFromCard(card))
		labelTop.TextStyle = fyne.TextStyle{Monospace: true}
		labelTop.Move(fyne.NewPos(CARD_LABEL_SPACING, 0))
		labelBottom = canvas.NewText(card.String(), getColorFromCard(card))
		labelBottom.TextStyle = fyne.TextStyle{Monospace: true}
	} else {
		rectangle = canvas.NewRectangle(BACK_COLOR)
		labelTop = canvas.NewText(card.String(), color.Transparent)
		labelBottom = canvas.NewText(card.String(), color.Transparent)
	}
	if horizontal {
		rectangle.Resize(fyne.NewSize(CARD_HEIGHT, CARD_WIDTH))
		labelBottom.Move(fyne.NewPos(CARD_HEIGHT-labelBottom.MinSize().Width-CARD_LABEL_SPACING, CARD_WIDTH-labelBottom.MinSize().Height))
	} else {
		rectangle.Resize(fyne.NewSize(CARD_WIDTH, CARD_HEIGHT))
		labelBottom.Move(fyne.NewPos(CARD_WIDTH-labelBottom.MinSize().Width-CARD_LABEL_SPACING, CARD_HEIGHT-labelBottom.MinSize().Height))
	}
	rectangle.StrokeWidth = 1
	rectangle.StrokeColor = color.Black

	return container.NewWithoutLayout(rectangle, labelTop, labelBottom)
}

func getColorFromCard(card *card) color.Color {
	switch {
	case HEARTS_SUIT == card.suit || DIAMONDS_SUIT == card.suit:
		return RED
	default:
		return color.Black
	}
}

type trickWidget struct {
	widget.BaseWidget
	cards       [4]*card
	firstPlayer uint8
}

func NewTrickWidget(firstPlayer uint8, cards [4]*card) *trickWidget {
	trickWidget := &trickWidget{}
	trickWidget.ExtendBaseWidget(trickWidget)
	trickWidget.SetTrickWidget(firstPlayer, cards)

	return trickWidget
}

func (p *trickWidget) SetTrickWidget(firstPlayer uint8, cards [4]*card) {
	p.firstPlayer = firstPlayer
	p.cards = cards
	p.Refresh()
}

type trickWidgetRenderer struct {
	t              *trickWidget
	background     *canvas.Rectangle
	paintableCards []fyne.CanvasObject
}

func (p *trickWidget) CreateRenderer() fyne.WidgetRenderer {
	return newTrickWidgetRenderer(p)
}

func newTrickWidgetRenderer(t *trickWidget) *trickWidgetRenderer {
	render := &trickWidgetRenderer{}
	render.t = t
	render.background = canvas.NewRectangle(color.Transparent)
	render.Refresh()

	return render
}

func (p *trickWidgetRenderer) Destroy() {}

func (p *trickWidgetRenderer) Layout(size fyne.Size) {
	p.background.Resize(size)
	offset := 0
	overlap := 35
	for i := range p.paintableCards {
		p.paintableCards[(i-int(p.t.firstPlayer)+5)%4].Move(fyne.NewPos(
			(size.Width-CARD_WIDTH)/2+float32(((i+1)%2)*(i%4-1)*(CARD_WIDTH/2-TRICK_CARD_OVERLAP))+float32((i%2)*(i%4-1)*CARD_COUNTERCLOCKWISE_OFFSET),
			(size.Height-CARD_HEIGHT)/2+float32((i%2)*(i%4-2)*(CARD_HEIGHT/2-TRICK_CARD_OVERLAP))+float32((((i+1)%2)*(i+1)%4)*-CARD_COUNTERCLOCKWISE_OFFSET),
		))
		offset += overlap
	}
}

func (p *trickWidgetRenderer) Refresh() {
	p.paintableCards = []fyne.CanvasObject{} //TODO: do stałej
	for i := range p.t.cards {
		if p.t.cards[(i+int(p.t.firstPlayer))%4] == nil { //TODO: do stałej
			invisibleCard := container.NewWithoutLayout()
			p.paintableCards = append(p.paintableCards, invisibleCard)
			continue
		}
		paintableCard := CreatePaintableCard(p.t.cards[(i+int(p.t.firstPlayer))%4], true, false) //TODO: 4 do stałej
		// paintableCard := NewCardWidget(v, true, false)
		p.paintableCards = append(p.paintableCards, paintableCard)
	}
	canvas.Refresh(p.t)
}

func (p *trickWidgetRenderer) MinSize() fyne.Size {
	side := CARD_HEIGHT*2 - 2*TRICK_CARD_OVERLAP

	return fyne.Size{Width: float32(side), Height: float32(side)}
}

func (p *trickWidgetRenderer) Objects() []fyne.CanvasObject {
	result := []fyne.CanvasObject{p.background}

	return append(result, p.paintableCards...)
}
