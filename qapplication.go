package qtwidgets

import (
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
	rvp := qtrt.Callany[voidptr](nil, nil, len(argv), argv, flags)
	return QApplicationFromptr(rvp)
}

func (me *QApplication) Exec() int {
	rvp := qtrt.Callany[int](nil, me)
	return rvp
}
func (me *QApplication) Exit(code int) {
	qtrt.Callany[int](nil, me)
}

func (me *QWidget) Show() {
	qtrt.Callany[int](nil, me)
}
func (me *QWidget) Hide() {
	qtrt.Callany[int](nil, me)
}

type QWidget struct {
	*qtrt.CObject
}

func QWidgetFromptr(ptr voidptr) *QWidget {
	me := &QWidget{qtrt.CObjectFromptr(ptr)}
	return me
}
func (me *QWidget) Dtor() {
	qtrt.Callany0(nil, me)
}

// size=8
type QSize struct {
	W int32
	H int32
}

func (me *QWidget) Size() QSize {
	rv := qtrt.Callany[QSize](nil, me)
	return rv
}
func (me *QWidget) Width() int {
	return qtrt.Callany[int](nil, me)
}
func (me *QWidget) Height() int {
	return qtrt.Callany[int](nil, me)
}
func (me *QWidget) X() int {
	return qtrt.Callany[int](nil, me)
}
func (me *QWidget) Y() int {
	return qtrt.Callany[int](nil, me)
}

type QMainWindow struct {
	*QWidget
}

func QMainWindowFromptr(ptr voidptr) *QMainWindow {
	return &QMainWindow{QWidgetFromptr(ptr)}
}
func (me *QMainWindow) Dtor() {
	qtrt.Callany0(nil, me)
}

func NewQMainWindow(parent *QWidget, flags int) *QMainWindow {
	// log.Println(qtclzsz.Get("QMainWindow"))
	cthis := qtrt.Callany[voidptr](nil, nil, parent, flags)
	return QMainWindowFromptr(cthis)
}
