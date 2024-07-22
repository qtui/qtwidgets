package qtwidgets

import (
	"github.com/qtui/qtrt"
)

type QApplication struct {
	qtrt.CObject
}

func QApplicationFromptr(ptr voidptr) *QApplication {
	me := &QApplication{qtrt.CObject{ptr}}
	return me
}

// QApplication::QApplication(int&, char**, int)
func NewQApplication(argv []string, flags int) *QApplication {
	rvp := qtrt.Callany(nil, len(argv), argv, flags)
	return QApplicationFromptr(rvp.Ptr())
}

func (me *QApplication) Exec() int {
	rvp := qtrt.Callany(me.GetCthis_())
	return rvp.Int()
}
