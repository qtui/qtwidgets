package qtwidgets

import (
	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
)

type QLayoutItem struct {
	*qtrt.CObject
}

type QLayoutItemITF interface {
	QLayoutItemPTR() *QLayoutItem
}

func (me *QLayoutItem) QLayoutItemPTR() *QLayoutItem { return me }

func QLayoutItemFromptr(ptr voidptr) *QLayoutItem {
	return &QLayoutItem{qtrt.CObjectFromptr(ptr)}
}

type QSpacerItem struct {
	*QLayoutItem
}

func QSpacerItemFromptr(ptr voidptr) *QSpacerItem {
	return &QSpacerItem{QLayoutItemFromptr(ptr)}
}

// weak symbol
func NewQSpacerItem(w, h int, wp, hp int) *QSpacerItem {
	rv := qtrt.Callany[voidptr](nil, w, h, wp, hp)
	return QSpacerItemFromptr(rv)
}

type QWidgetItem struct {
	*QLayoutItem
}

func QWidgetItemFromptr(ptr voidptr) *QWidgetItem {
	return &QWidgetItem{QLayoutItemFromptr(ptr)}
}

type QWidgetItemV2 struct {
	*QWidgetItem
}

func QWidgetItemV2Fromptr(ptr voidptr) *QWidgetItemV2 {
	return &QWidgetItemV2{QWidgetItemFromptr(ptr)}
}

type QLayout struct {
	*QLayoutItem
}

type QLayoutITF interface {
	QLayoutPTR() *QLayout
}

func (me *QLayout) QLayoutPTR() *QLayout { return me }

func QLayoutFromptr(ptr voidptr) *QLayout {
	return &QLayout{QLayoutItemFromptr(ptr)}
}

func (me *QLayout) AddWidget(w QWidgetITF) {
	qtrt.Callany0(me, w)
}
func (me *QLayout) AddItem(item QLayoutItemITF) {
	qtrt.Callany0(me, item)
}
func (me *QLayout) RemoveWidget(w QWidgetITF) {
	qtrt.Callany0(me, w)
}
func (me *QLayout) RemoveItem(item QLayoutItemITF) {
	qtrt.Callany0(me, item)
}

type QBoxLayout struct {
	*QLayout
}

func QBoxLayoutFromptr(ptr voidptr) *QBoxLayout {
	return &QBoxLayout{QLayoutFromptr(ptr)}
}

// 这个方法不能传递QBoxLayout??? crash
func (me *QBoxLayout) AddItem(item QLayoutItemITF) {
	qtrt.Callany0(me, item)
}
func (me *QBoxLayout) AddLayout(item QLayoutITF, stretch ...int) {
	qtrt.Callany0(me, item, gopp.FirstofGv(stretch))
}

func (me *QBoxLayout) Direction() (dir int) {
	dir = qtrt.Callany[int](me)
	return
}
func (me *QBoxLayout) SetDirection(dir int) {
	qtrt.Callany0(me, dir)
}

type QHBoxLayout struct {
	*QBoxLayout
}

func QHBoxLayoutFromptr(ptr voidptr) *QHBoxLayout {
	return &QHBoxLayout{QBoxLayoutFromptr(ptr)}
}

func NewQHBoxLayout(parent QWidgetITF) *QHBoxLayout {
	rv := qtrt.Callany[voidptr](nil, parent)
	return QHBoxLayoutFromptr(rv)
}

type QVBoxLayout struct {
	*QBoxLayout
}

func QVBoxLayoutFromptr(ptr voidptr) *QVBoxLayout {
	return &QVBoxLayout{QBoxLayoutFromptr(ptr)}
}
func NewQVBoxLayout(parent QWidgetITF) *QVBoxLayout {
	rv := qtrt.Callany[voidptr](nil, parent)
	return QVBoxLayoutFromptr(rv)
}

type QGridLayout struct {
	*QLayout
}

func QGridLayoutFromptr(ptr voidptr) *QGridLayout {
	return &QGridLayout{QLayoutFromptr(ptr)}
}

type QFormLayout struct {
	*QLayout
}

func QFormLayoutFromptr(ptr voidptr) *QFormLayout {
	return &QFormLayout{QLayoutFromptr(ptr)}
}
