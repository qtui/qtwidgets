package qtwidgets

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QApplication struct {
	*qtcore.QGuiApplication
}

func QApplicationFromptr(ptr voidptr) *QApplication {
	me := &QApplication{qtcore.QGuiApplicationFromptr(ptr)}
	return me
}

func QApp() *QApplication { return (*QApplication)(nil).Instance() }
func (me *QApplication) Instance() *QApplication {
	return QApplicationFromptr(qtcore.QCoreApplication_Instance().GetCthis())
}

// QApplication::QApplication(int&, char**, int)
func NewQApplication(argv []string, flags int) *QApplication {
	rvp := qtrt.Callany[voidptr](nil, len(argv), argv, flags)
	return QApplicationFromptr(rvp)
}

func (me *QApplication) Exec() int {
	rvp := qtrt.Callany[int](me)
	return rvp
}
func (me *QApplication) Exit(code int) {
	qtrt.Callany0(me)
}
func (me *QApplication) SetStyleSheet(s string) {
	qtrt.Callany0(me, s)
}
func (me *QApplication) AboutQt() {
	QApplication_aboutQt()
}
func QApplication_aboutQt() { qtrt.Callany0(nil) }

func (me *QWidget) Show() {
	qtrt.Callany0(me)
}
func (me *QWidget) Hide() {
	qtrt.Callany0(me)
}

type QWidget struct {
	*qtcore.QObject
}
type QWidgetITF interface {
	QWidgetPTR() *QWidget
}

func (me *QWidget) QWidgetPTR() *QWidget { return me }

func QWidgetFromptr(ptr voidptr) *QWidget {
	me := &QWidget{qtcore.QObjectFromptr(ptr)}
	return me
}

func (me *QWidget) Dtor() {
	qtrt.Callany0(me)
}

func NewQWidget(parent QWidgetITF, f ...int) *QWidget {
	var f0 = gopp.FirstofGv(f)
	rv := qtrt.Callany[voidptr](nil, parent, f0)
	return QWidgetFromptr(rv)
}
func (me *QWidget) SetLayout(lo QLayoutITF) {
	qtrt.Callany0(me, lo)
}
func (me *QWidget) Resize(w, h int)            { qtrt.Callany0(me, w, h) }
func (me *QWidget) SetVisible(b bool)          { qtrt.Callany0(me, b) }
func (me *QWidget) SetHidden(b bool)           { qtrt.Callany0(me, b) }
func (me *QWidget) IsVisible() bool            { return qtrt.Callany[bool](me) }
func (me *QWidget) SetWindowOpacity(v float64) { qtrt.Callany0(me, v) }
func (me *QWidget) SetToolTip(v string)        { qtrt.Callany0(me, v) }
func (me *QWidget) SetStatusTip(v string)      { qtrt.Callany0(me, v) }

// size=8
type QSize struct {
	W int32
	H int32
}

func (me *QWidget) Size() QSize {
	rv := qtrt.Callany[QSize](me)
	return rv
}
func (me *QWidget) Width() int {
	return qtrt.Callany[int](me)
}
func (me *QWidget) Height() int {
	return qtrt.Callany[int](me)
}
func (me *QWidget) X() int {
	return qtrt.Callany[int](me)
}
func (me *QWidget) Y() int {
	return qtrt.Callany[int](me)
}

type QMainWindow struct {
	*QWidget
}

func QMainWindowFromptr(ptr voidptr) *QMainWindow {
	return &QMainWindow{QWidgetFromptr(ptr)}
}
func (me *QMainWindow) Dtor() {
	qtrt.Callany0(me)
}

func NewQMainWindow(parent QWidgetITF, flags int) *QMainWindow {
	// log.Println(qtclzsz.Get("QMainWindow"))
	cthis := qtrt.Callany[voidptr](nil, parent, flags)
	return QMainWindowFromptr(cthis)
}

func (me *QMainWindow) CentralWidget() (w *QWidget) {
	rv := qtrt.Callany[voidptr](me)
	w = QWidgetFromptr(rv)
	return
}

func (me *QMainWindow) SetCentralWidget(w QWidgetITF) {
	qtrt.Callany0(me, w)
}

type QToolBar struct {
	*QWidget
}

type QStatusBar struct {
}
