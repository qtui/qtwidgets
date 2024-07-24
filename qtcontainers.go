package qtwidgets

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
)

type QFrame struct {
	*QWidget
}

func QFrameFromptr(ptr voidptr) *QFrame {
	return &QFrame{QWidgetFromptr(ptr)}
}

type QStackedWidget struct {
	*QFrame
}

func QStackedWidgetFromptr(ptr voidptr) *QStackedWidget {
	return &QStackedWidget{QFrameFromptr(ptr)}
}
func NewQStackedWidget(parent ...*QWidget) *QStackedWidget {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QStackedWidgetFromptr(rv)
}

func (me *QStackedWidget) AddWidget(w *QWidget) int {
	rv := qtrt.Callany[int](me, w)
	return rv
}
func (me *QStackedWidget) InsertWidget(idx int, w *QWidget) int {
	rv := qtrt.Callany[int](me, idx, w)
	return rv
}
func (me *QStackedWidget) RemoveWidget(w *QWidget) {
	qtrt.Callany0(me, w)
}
func (me *QStackedWidget) CurrentWidget() *QWidget {
	rv := qtrt.Callany[voidptr](me)
	return QWidgetFromptr(rv)
}
func (me *QStackedWidget) CurrentIndex() int {
	rv := qtrt.Callany[int](me)
	return rv
}
func (me *QStackedWidget) SetCurrentIndex(idx int) {
	qtrt.Callany[int](me, idx)
}
func (me *QStackedWidget) Count() int {
	rv := qtrt.Callany[int](me)
	return rv
}

type QTabWidget struct {
	*QWidget
}

func QTabWidgetFromptr(ptr voidptr) *QTabWidget {
	return &QTabWidget{QWidgetFromptr(ptr)}
}

type QSplitter struct {
	*QFrame
}

func QSplitterFromptr(ptr voidptr) *QSplitter {
	return &QSplitter{QFrameFromptr(ptr)}
}

func NewQSplitter(parent *QWidget) *QSplitter {
	rv := qtrt.Callany[voidptr](nil, parent)
	return QSplitterFromptr(rv)
}

func (me *QSplitter) AddWidget(w *QWidget) {
	qtrt.Callany0(me, w)
}
func (me *QSplitter) InsertWidget(idx int, w *QWidget) {
	qtrt.Callany0(me, idx, w)
}
func (me *QSplitter) ReplaceWidget(idx int, w *QWidget) (wret *QWidget) {
	rv := qtrt.Callany[voidptr](me, idx, w)
	wret = QWidgetFromptr(rv)
	return
}
