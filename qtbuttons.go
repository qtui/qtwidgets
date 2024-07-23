package qtwidgets

import (
	"log"

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
	qtrt.Callany[int](nil, me, text)
}

type QPushButton struct {
	*QAbstractButton
}

func QPushButtonFromptr(ptr voidptr) *QPushButton {
	me := &QPushButton{QAbstractButtonFromptr(ptr)}
	return me
}

func NewQPushButton() *QPushButton {
	rvp := qtrt.Callany[voidptr](nil, nil, nil)
	log.Println(rvp)

	return QPushButtonFromptr(rvp)
}

func (me *QPushButton) SetFlat(b bool) {
	qtrt.Callany[int](nil, me, b)
}
