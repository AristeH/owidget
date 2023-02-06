package otable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tappableIcon struct {
	widget.Icon
	t *OTable
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (t *tappableIcon) Tapped(ev *fyne.PointEvent) {
}
func (t *tappableIcon) KeyDown(key *fyne.KeyEvent) {
}

// TypedKey  fyne.Focusable
func (t *tappableIcon) TypedKey(ev *fyne.KeyEvent) {
	t.t.TypedKey(ev)
}

// FocusGained  fyne.Focusable
func (t *tappableIcon) FocusGained() {
}

func (t *tappableIcon) TypedRune(r rune) {
}

// FocusLost  fyne.Focusable
func (t *tappableIcon) FocusLost() {

}
