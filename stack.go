package confish

import (
	"container/list"
)

type confItem struct {
	Elem interface{}
}

func newConfItem(elem interface{}) *confItem {
	return &confItem{
		Elem: elem,
	}
}

type confStack struct {
	list *list.List
}

func newConfStack() *confStack {
	return &confStack{list.New()}
}

func (cs *confStack) Push(cfgVar interface{}) {
	item := newConfItem(cfgVar)
	cs.list.PushBack(item)
}

func (cs *confStack) Pop() (*confItem, bool) {
	if cs.list.Len() == 0 {
		return nil, false
	}

	element := cs.list.Back()
	cs.list.Remove(element)

	return element.Value.(*confItem), true
}

func (cs *confStack) Peek() *confItem {
	if cs.list.Len() == 0 {
		return nil
	}

	element := cs.list.Back()
	return element.Value.(*confItem)
}

func (cs *confStack) Size() int {
	return cs.list.Len()
}

func (cs *confStack) IsEmpty() bool {
	return (cs.list.Len() == 0)
}
