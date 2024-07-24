package qtwidgets

import "github.com/qtui/qtcore"

type QLabel struct {
	*QFrame
}

func QLabelFromptr(ptr voidptr) *QLabel {
	return &QLabel{QFrameFromptr(ptr)}
}

type QComboBox struct {
	*QWidget
}

func QComboBoxFromptr(ptr voidptr) *QComboBox {
	return &QComboBox{QWidgetFromptr(ptr)}
}

type QCompleter struct {
	*qtcore.QObject
}

func QCompleterFromptr(ptr voidptr) *QCompleter {
	return &QCompleter{qtcore.QObjectFromptr(ptr)}
}

type QLineEdit struct {
	*QWidget
}

func QLineEditFromptr(ptr voidptr) *QLineEdit {
	return &QLineEdit{QWidgetFromptr(ptr)}
}

type QTextEdit struct {
	*QWidget
}

func QTextEditFromptr(ptr voidptr) *QTextEdit {
	return &QTextEdit{QWidgetFromptr(ptr)}
}

type QPlainTextEdit struct {
	*QWidget
}

func QPlainTextEditFromptr(ptr voidptr) *QPlainTextEdit {
	return &QPlainTextEdit{QWidgetFromptr(ptr)}
}

type QTextBrowser struct {
	*QWidget
}

func QTextBrowserFromptr(ptr voidptr) *QTextBrowser {
	return &QTextBrowser{QWidgetFromptr(ptr)}
}

type QCheckBox struct {
	*QWidget
}

func QCheckBoxFromptr(ptr voidptr) *QCheckBox {
	return &QCheckBox{QWidgetFromptr(ptr)}
}

type QGroupBox struct {
	*QWidget
}

func QGroupBoxFromptr(ptr voidptr) *QGroupBox {
	return &QGroupBox{QWidgetFromptr(ptr)}
}

type QProgressBar struct {
	*QWidget
}

func QProgressBarFromptr(ptr voidptr) *QProgressBar {
	return &QProgressBar{QWidgetFromptr(ptr)}
}

type QToolBox struct {
	*QWidget
}

func QToolBoxFromptr(ptr voidptr) *QToolBox {
	return &QToolBox{QWidgetFromptr(ptr)}
}

type QToolTip struct {
	*QWidget
}

func QToolTipFromptr(ptr voidptr) *QToolTip {
	return &QToolTip{QWidgetFromptr(ptr)}
}
