package prlist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	pl := New()

	pl.Push(1, 2)
	pl.Push(2, 4)
	pl.Push(3, 2)
	pl.Push(4)
	pl.Push(5, 5)
	e := pl.Push(6, 4)
	pl.Push(7, 5)
	pl.Push(8)

	pl.test()
	for e := pl.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
	fmt.Println()
	t.Log(pl.Remove(e))
	pl.Push(9, 4)
	fmt.Println()

	pl2 := New()
	pl2.PushList(pl)
	pl.PushList(pl)

	//t.Log(pl.Len())
	for i := pl.Len(); i > 0; i-- {
		t.Log(pl.Pop())
	}
	pl.test()
	pl2.test()
	for i := pl2.Len(); i > 0; i-- {
		t.Log(pl2.Pop())
	}
}

func (pl *PrList) test() {
	fmt.Println(pl.l.Len(), pl.gl.Len())
}
