package qtwidgets

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtrt"
)

type QLabel struct {
	*QFrame
}

func QLabelFromptr(ptr voidptr) *QLabel {
	return &QLabel{QFrameFromptr(ptr)}
}

func NewQLabel(text string, parent QWidgetITF, f ...int) *QLabel {
	rv := qtrt.Callany[voidptr](nil, text, parent, gopp.FirstofGv(f))
	return QLabelFromptr(rv)
}
func (me *QLabel) SetText(text string) { qtrt.Callany0(me, text) }
func (me *QLabel) SetTextInteractionFlags(f int) {
	qtrt.Callany0(me, f)
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

func NewQTextEdit(parent ...QWidgetITF) *QTextEdit {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QTextEditFromptr(rv)
}
func (me *QTextEdit) SetReadOnly(b bool) {
	qtrt.Callany0(me, b)
}
func (me *QTextEdit) IsReadOnly() (b bool) {
	b = qtrt.Callany[bool](me)
	return
}

func (me *QTextEdit) SetPlainText(text string) {
	qtrt.Callany0(me, text)
}
func (me *QTextEdit) SetHtml(text string) {
	qtrt.Callany0(me, text)
}
func (me *QTextEdit) SetMarkdown(text string) {
	qtrt.Callany0(me, text)
}
func (me *QTextEdit) SetText(text string) {
	qtrt.Callany0(me, text)
}
func (me *QTextEdit) Append(text string) {
	qtrt.Callany0(me, text)
}

type QPlainTextEdit struct {
	*QWidget
}

func QPlainTextEditFromptr(ptr voidptr) *QPlainTextEdit {
	return &QPlainTextEdit{QWidgetFromptr(ptr)}
}

// 为什么是从Edit继承来的
type QTextBrowser struct {
	*QTextEdit
}

func QTextBrowserFromptr(ptr voidptr) *QTextBrowser {
	return &QTextBrowser{QTextEditFromptr(ptr)}
}

func NewQTextBrowser(parent ...QWidgetITF) *QTextBrowser {
	rv := qtrt.Callany[voidptr](nil, gopp.FirstofGv(parent))
	return QTextBrowserFromptr(rv)
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
