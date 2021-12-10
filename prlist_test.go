package prlist

import (
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
