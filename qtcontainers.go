package qtwidgets

type QStackedWidget struct {
	*QWidget
}

func QStackedWidgetFromptr(ptr voidptr) *QStackedWidget {
	return &QStackedWidget{QWidgetFromptr(ptr)}
}

type QTabWidget struct {
	*QWidget
}

func QTabWidgetFromptr(ptr voidptr) *QTabWidget {
	return &QTabWidget{QWidgetFromptr(ptr)}
}

type QFrame struct {
	*QWidget
}

func QFrameFromptr(ptr voidptr) *QFrame {
	return &QFrame{QWidgetFromptr(ptr)}
}

type QSplitter struct {
	*QFrame
}

func QSplitterFromptr(ptr voidptr) *QSplitter {
	return &QSplitter{QFrameFromptr(ptr)}
}
