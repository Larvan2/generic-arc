package main

import (
    "container/list"
)

type entry struct {
    key interface{}
    value interface{}
    ll *list.List
    el *list.Element
}

func (e *entry) setLRU(list *list.List) {
    e.ll = list
    e.el = e.ll.PushBack(e)
}

func (e *entry) setMRU(list *list.List) {
    e.ll = list
    e.el = e.ll.PushFront(e)
}

func (e *entry) detach() {
    if e.ll != nil {
        e.ll.Remove(e.el)
    }
}