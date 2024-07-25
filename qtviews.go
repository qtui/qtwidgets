package qtwidgets

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
)

type QAbstractScrollArea struct {
	*QFrame
}

func QAbstractScrollAreaFromptr(ptr voidptr) *QAbstractScrollArea {
	return &QAbstractScrollArea{QFrameFromptr(ptr)}
}

type QAbstractItemView struct {
	*QAbstractScrollArea
}

func QAbstractItemViewFromptr(ptr voidptr) *QAbstractItemView {
	return &QAbstractItemView{QAbstractScrollAreaFromptr(ptr)}
}

type QTableView struct {
	*QAbstractItemView
}

func QTableViewFromptr(ptr voidptr) *QTableView {
	return &QTableView{QAbstractItemViewFromptr(ptr)}
}

type QTableWidgetItem struct {
	*qtrt.CObject
}

func QTableWidgetItemFromptr(ptr voidptr) *QTableWidgetItem {
	return &QTableWidgetItem{qtrt.CObjectFromptr(ptr)}
}

func NewQTableWidgetItem(text string, typ ...int) *QTableWidgetItem {
	rv := qtrt.Callany[voidptr](nil, text, gopp.FirstofGv(typ))
	return QTableWidgetItemFromptr(rv)
}

type QTableWidget struct {
	*QTableView
}

func QTableWidgetFromptr(ptr voidptr) *QTableWidget {
	return &QTableWidget{QTableViewFromptr(ptr)}
}
func NewQTableWidget(parent ...QWidgetITF) *QTableWidget {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QTableWidgetFromptr(rv)
}

func (me *QTableWidget) SetRowCount(c int) {
	qtrt.Callany0(me, c)
}
func (me *QTableWidget) SetColumnCount(c int) {
	qtrt.Callany0(me, c)
}

func (me *QTableWidget) SetItem(row, col int, item *QTableWidgetItem) {
	qtrt.Callany0(me, row, col, item)
}
func (me *QTableWidget) InsertRow(row int) { qtrt.Callany0(me, row) }

type QListView struct {
	*QAbstractItemView
}

func QListViewFromptr(ptr voidptr) *QListView {
	return &QListView{QAbstractItemViewFromptr(ptr)}
}

type QListWidget struct {
	*QListView
}

func QListWidgetFromptr(ptr voidptr) *QListWidget {
	return &QListWidget{QListViewFromptr(ptr)}
}

type QTreeView struct {
	*QAbstractItemView
}

func QTreeViewFromptr(ptr voidptr) *QTreeView {
	return &QTreeView{QAbstractItemViewFromptr(ptr)}
}

type QTreeWidget struct {
	*QTreeView
}

func QTreeWidgetFromptr(ptr voidptr) *QTreeWidget {
	return &QTreeWidget{QTreeViewFromptr(ptr)}
}
