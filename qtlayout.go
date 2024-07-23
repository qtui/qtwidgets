package qtwidgets

import "github.com/qtui/qtrt"

type QLayoutItem struct {
	*qtrt.CObject
}

func QLayoutItemFromptr(ptr voidptr) *QLayoutItem {
	return &QLayoutItem{qtrt.CObjectFromptr(ptr)}
}

type QSpacerItem struct {
	*QLayoutItem
}

func QSpacerItemFromptr(ptr voidptr) *QSpacerItem {
	return &QSpacerItem{QLayoutItemFromptr(ptr)}
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

func QLayoutFromptr(ptr voidptr) *QLayout {
	return &QLayout{QLayoutItemFromptr(ptr)}
}

type QBoxLayout struct {
	*QLayout
}

func QBoxLayoutFromptr(ptr voidptr) *QBoxLayout {
	return &QBoxLayout{QLayoutFromptr(ptr)}
}

type QHBoxLayout struct {
	*QBoxLayout
}

func QHBoxLayoutFromptr(ptr voidptr) *QHBoxLayout {
	return &QHBoxLayout{QBoxLayoutFromptr(ptr)}
}

type QVBoxLayout struct {
	*QBoxLayout
}

func QVBoxLayoutFromptr(ptr voidptr) *QVBoxLayout {
	return &QVBoxLayout{QBoxLayoutFromptr(ptr)}
}

type QGridLayout struct {
	*QLayout
}

func QGridLayoutFromptr(ptr voidptr) *QGridLayout {
	return &QGridLayout{QLayoutFromptr(ptr)}
}
