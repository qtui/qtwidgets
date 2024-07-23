package qtwidgets

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

type QTableWidget struct {
	*QTableView
}

func QTableWidgetFromptr(ptr voidptr) *QTableWidget {
	return &QTableWidget{QTableViewFromptr(ptr)}
}

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
