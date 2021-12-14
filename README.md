# prlist

priority list, 优先级链表

## Index

```go
type Element 
func (e *Element) Prev() *Element
func (e *Element) Next() *Element 

type PrList 
func New() *PrList 
func (pl *PrList) Len() int 
func (pl *PrList) Front() *Element
func (pl *PrList) Back() *Element 
func (pl *PrList) Remove(e *Element) interface{} 
func (pl *PrList) Push(v interface{}, priority ...uint32) *Element
func (pl *PrList) PushList(other *PrList)
func (pl *PrList) Pop() interface{}

```

## Usage

```go
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

// output
//    prlist_test.go:23: 5
//    prlist_test.go:23: 7
//    prlist_test.go:23: 2
//    prlist_test.go:23: 6
//    prlist_test.go:23: 1
//    prlist_test.go:23: 3
//    prlist_test.go:23: 4
//    prlist_test.go:23: 8
//
//    prlist_test.go:28: 6
//
//    prlist_test.go:39: 5
//    prlist_test.go:39: 7
//    prlist_test.go:39: 2
//    prlist_test.go:39: 9
//    prlist_test.go:39: 1
//    prlist_test.go:39: 3
//    prlist_test.go:39: 4
//    prlist_test.go:39: 8
//
//    prlist_test.go:43: 5
//    prlist_test.go:43: 7
//    prlist_test.go:43: 2
//    prlist_test.go:43: 9
//    prlist_test.go:43: 1
//    prlist_test.go:43: 3
//    prlist_test.go:43: 4
//    prlist_test.go:43: 8
```