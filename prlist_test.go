package prlist

import (
	"container/list"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	pl := New()

	// push
	pl.Push(1, 2)
	pl.Push(2, 4)
	pl.Push(3, 2)
	pl.Push(4)
	pl.Push(5, 5)
	e := pl.Push(6, 4)
	pl.Push(7, 5)
	pl.Push(8)

	// range
	for e := pl.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}

	fmt.Println()
	// remove
	t.Log(pl.Remove(e))
	pl.Push(9, 4)

	fmt.Println()
	// push list
	pl2 := New()
	pl2.PushList(pl)
	pl.PushList(pl)

	// pop
	for i := pl.Len(); i > 0; i-- {
		t.Log(pl.Pop())
	}
	fmt.Println()
	for i := pl2.Len(); i > 0; i-- {
		t.Log(pl2.Pop())
	}
}

func TestList(t *testing.T) {
	l := list.New()
	t.Logf("len:%d list:%v \n", l.Len(), l)
	e1 := l.PushBack(1)
	t.Logf("len:%d list:%v \n", l.Len(), l)
	l.Init()
	t.Logf("len:%d list:%v \n", l.Len(), l)
	l.Remove(e1)
	t.Logf("len:%d list:%v \n", l.Len(), l)
	l.PushBack(2)
	t.Logf("len:%d list:%v \n", l.Len(), l)
	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}
