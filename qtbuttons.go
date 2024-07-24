package qtwidgets

import (
	"log"

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

func NewQPushButtonz0(parent ...*QWidget) *QPushButton {
	var p = gopp.FirstofGv(parent)
	rvp := qtrt.Callany[voidptr](nil, p)
	log.Println(rvp)

	return QPushButtonFromptr(rvp)
}
func NewQPushButton(text string, parent ...*QWidget) *QPushButton {
	var p = gopp.FirstofGv(parent)
	rvp := qtrt.Callany[voidptr](nil, text, p)
	log.Println(rvp)

	return QPushButtonFromptr(rvp)
}

func (me *QPushButton) SetFlat(b bool) {
	qtrt.Callany0(me, b)
}
