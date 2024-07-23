package qtwidgets

import (
	"log"

	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QApplication struct {
	*qtcore.QObject
}

func QApplicationFromptr(ptr voidptr) *QApplication {
	me := &QApplication{qtcore.QObjectFromptr(ptr)}
	return me
}

// QApplication::QApplication(int&, char**, int)
func NewQApplication(argv []string, flags int) *QApplication {
	rvp := qtrt.Callany(nil, nil, len(argv), argv, flags)
	return QApplicationFromptr(rvp.Ptr())
}

func (me *QApplication) Exec() int {
	rvp := qtrt.Callany(nil, me.GetCthis())
	return rvp.Int()
}
func (me *QApplication) Exit(code int) {
	qtrt.Callany(nil, me.GetCthis())
}

type QAbstractButton struct {
	*QWidget
}

func QAbstractButtonFromptr(ptr voidptr) *QAbstractButton {
	me := &QAbstractButton{QWidgetFromptr(ptr)}
	return me
}

func (me *QAbstractButton) SetText(text string) {
	qtrt.Callany(nil, me.GetCthis(), text)
}

type QPushButton struct {
	*QAbstractButton
}

func QPushButtonFromptr(ptr voidptr) *QPushButton {
	me := &QPushButton{QAbstractButtonFromptr(ptr)}
	return me
}

func NewQPushButton() *QPushButton {
	rvp := qtrt.Callany(nil, nil, nil)
	log.Println(rvp.Ptr())

	return QPushButtonFromptr(rvp.Ptr())
}

func (me *QPushButton) SetFlat(b bool) {
	qtrt.Callany(nil, me.GetCthis(), b)
}

func (me *QWidget) Show() {
	qtrt.Callany(nil, me.GetCthis())
}

type QWidget struct {
	*qtrt.CObject
}

func QWidgetFromptr(ptr voidptr) *QWidget {
	me := &QWidget{qtrt.CObjectFromptr(ptr)}
	return me
}
