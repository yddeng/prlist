package prlist

import (
	"container/list"
)

// Element is an element of a linked prlist.
type Element struct {
	// The value stored with this element.
	Value interface{}

	e  *list.Element
	g  *guard
	pl *PrList
}

// Prev returns the previous prlist element or nil.
func (e *Element) Prev() *Element {
	if e.e.Prev() == nil {
		return nil
	}
	return e.e.Prev().Value.(*Element)
}

// Next returns the next prlist element or nil.
func (e *Element) Next() *Element {
	if e.e.Next() == nil {
		return nil
	}
	return e.e.Next().Value.(*Element)
}

type guard struct {
	priority uint32
	e        *list.Element // pl.gl.Element
	mark     *list.Element // pl.l.Element
}

type PrList struct {
	l  *list.List // element
	gl *list.List // guard
}

// New returns an initialized prlist.
func New() *PrList {
	return &PrList{
		l:  list.New(),
		gl: list.New(),
	}
}

// Init initializes or clears prlist pl.
func (pl *PrList) Init() *PrList {
	pl.l.Init()
	pl.gl.Init()
	return pl
}

// Len returns the number of elements of prlist pl.
// The complexity is O(1).
func (pl *PrList) Len() int {
	return pl.l.Len()
}

// Front returns the first element of prlist pl or nil if the prlist is empty.
func (pl *PrList) Front() *Element {
	if pl.l.Len() == 0 {
		return nil
	}
	return pl.l.Front().Value.(*Element)
}

// Back returns the last element of prlist pl or nil if the prlist is empty.
func (pl *PrList) Back() *Element {
	if pl.l.Len() == 0 {
		return nil
	}
	return pl.l.Back().Value.(*Element)
}

// Remove removes e from pl if e is an element of prlist pl.
// It returns the element value e.Value.
// The element must not be nil.
func (pl *PrList) Remove(e *Element) interface{} {
	if e.pl == pl {
		// guard.mark == e.e,
		if e.g.mark == e.e {
			prev := e.Prev()
			if prev == nil || prev.g != e.g {
				pl.gl.Remove(e.g.e)
			} else {
				e.g.mark = prev.e
			}
		}
		pl.l.Remove(e.e)
		e.pl = nil
	}
	return e.Value
}

// Push insert v with priority and return a new element.
func (pl *PrList) Push(v interface{}, priority ...uint32) *Element {
	pr := uint32(0)
	if len(priority) > 0 {
		pr = priority[0]
	}

	// find guard with priority and prev
	var g *guard
	var prev *list.Element
	for e := pl.gl.Front(); e != nil; e = e.Next() {
		if e.Value.(*guard).priority == pr {
			g = e.Value.(*guard)
			break
		} else if e.Value.(*guard).priority < pr {
			break
		}
		prev = e
	}

	// insert guard
	if g == nil {
		g = &guard{priority: pr}
		if prev == nil {
			g.e = pl.gl.PushFront(g)
		} else {
			g.mark = prev.Value.(*guard).mark
			g.e = pl.gl.InsertAfter(g, prev)
		}
	}

	// insert element
	elem := &Element{Value: v, pl: pl, g: g}
	if g.mark == nil {
		elem.e = pl.l.PushFront(elem)
	} else {
		elem.e = pl.l.InsertAfter(elem, g.mark)
	}
	g.mark = elem.e
	return elem
}

// PushList inserts a copy of another prlist.
func (pl *PrList) PushList(other *PrList) {
	if other == pl {
		return
	}
	for e := other.Front(); e != nil; e = e.Next() {
		pl.Push(e.Value, e.g.priority)
	}
}

// Pop returns the first element.Value of prlist pl and remove it.
func (pl *PrList) Pop() interface{} {
	e := pl.l.Front()
	if e != nil {
		return pl.Remove(e.Value.(*Element))
	} else {
		return nil
	}
}
