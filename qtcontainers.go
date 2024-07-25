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
func NewQStackedWidget(parent ...QWidgetITF) *QStackedWidget {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QStackedWidgetFromptr(rv)
}

func (me *QStackedWidget) AddWidget(w QWidgetITF) int {
	rv := qtrt.Callany[int](me, w)
	return rv
}
func (me *QStackedWidget) InsertWidget(idx int, w QWidgetITF) int {
	rv := qtrt.Callany[int](me, idx, w)
	return rv
}
func (me *QStackedWidget) RemoveWidget(w QWidgetITF) {
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
func (me *QStackedWidget) SetCurrentWidget(w QWidgetITF) {
	qtrt.Callany[int](me, w)
}
func (me *QStackedWidget) IndexOf(w QWidgetITF) int {
	rv := qtrt.Callany[int](me, w)
	return rv
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
func NewQTabWidget(parent ...QWidgetITF) *QTabWidget {
	rv := qtrt.Callany[voidptr](nil, gopp.LastofGv(parent))
	return QTabWidgetFromptr(rv)
}
func (me *QTabWidget) AddTab(w QWidgetITF, name string) int {
	rv := qtrt.Callany[int](me, w, name)
	return rv
}
func (me *QTabWidget) InsertTab(idx int, w QWidgetITF, name string) int {
	rv := qtrt.Callany[int](me, idx, w, name)
	return rv
}
func (me *QTabWidget) RemoveTab(idx int) {
	qtrt.Callany0(me, idx)
}
func (me *QTabWidget) CurrentIndex() int {
	rv := qtrt.Callany[int](me)
	return rv
}
func (me *QTabWidget) SetCurrentIndex(idx int) {
	qtrt.Callany[int](me, idx)
}
func (me *QTabWidget) SetCurrentWidget(w QWidgetITF) {
	qtrt.Callany[int](me, w)
}

func (me *QTabWidget) Count() int {
	rv := qtrt.Callany[int](me)
	return rv
}

type QSplitter struct {
	*QFrame
}

func QSplitterFromptr(ptr voidptr) *QSplitter {
	return &QSplitter{QFrameFromptr(ptr)}
}

func NewQSplitter(parent QWidgetITF) *QSplitter {
	rv := qtrt.Callany[voidptr](nil, parent)
	return QSplitterFromptr(rv)
}

func (me *QSplitter) AddWidget(w QWidgetITF) {
	qtrt.Callany0(me, w)
}
func (me *QSplitter) InsertWidget(idx int, w QWidgetITF) {
	qtrt.Callany0(me, idx, w)
}
func (me *QSplitter) ReplaceWidget(idx int, w QWidgetITF) (wret *QWidget) {
	rv := qtrt.Callany[voidptr](me, idx, w)
	wret = QWidgetFromptr(rv)
	return
}

func (me *QSplitter) SetOrientation(dir int) {
	qtrt.Callany0(me, dir)
}

func (me *QSplitter) SetStretchFactor(idx int, strech int) {
	qtrt.Callany0(me, idx, strech)
}
