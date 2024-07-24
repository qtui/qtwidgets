package qtwidgets

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
)

type QAbstractButton struct {
	*QWidget
}

func QAbstractButtonFromptr(ptr voidptr) *QAbstractButton {
	me := &QAbstractButton{QWidgetFromptr(ptr)}
	return me
}

func (me *QAbstractButton) SetText(text string) {
	qtrt.Callany0(me, text)
}

type QPushButton struct {
	*QAbstractButton
}

func QPushButtonFromptr(ptr voidptr) *QPushButton {
	me := &QPushButton{QAbstractButtonFromptr(ptr)}
	return me
}

func NewQPushButtonz0(parent ...QWidgetITF) *QPushButton {
	var p = gopp.FirstofGv(parent)
	rvp := qtrt.Callany[voidptr](nil, p)

	return QPushButtonFromptr(rvp)
}
func NewQPushButton(text string, parent ...QWidgetITF) *QPushButton {
	var p = gopp.FirstofGv(parent)
	rvp := qtrt.Callany[voidptr](nil, text, p)

	return QPushButtonFromptr(rvp)
}

func (me *QPushButton) SetFlat(b bool) {
	qtrt.Callany0(me, b)
}

type QToolButton struct {
	*QAbstractButton
}

func QToolButtonFromptr(ptr voidptr) *QToolButton {
	me := &QToolButton{QAbstractButtonFromptr(ptr)}
	return me
}
