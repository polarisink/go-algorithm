package skipList

import "unsafe"

type Element struct {
	elementHeader
	Value interface{}
	key   interface{}
	score float64

	prev         *Element
	prevTopLevel *Element
	list         *SkipList
}

type elementHeader struct {
	levels []*Element
}

func (header *elementHeader) Element() *Element {
	return (*Element)(unsafe.Pointer(header))
}

func newElement(list *SkipList, level int, score float64, key, value interface{}) *Element {
	return &Element{
		elementHeader: elementHeader{
			levels: make([]*Element, level),
		},
		Value: value,
		key:   key,
		score: score,
		list:  list,
	}
}

func (elem *Element) Next() *Element {
	if len(elem.levels) == 0 {
		return nil
	}
	return elem.levels[0]
}

func (elem *Element) Prev() *Element {
	return elem.prev
}

func (elem *Element) NextLevel(level int) *Element {
	if level < 0 || level >= len(elem.levels) {
		return nil
	}
	return elem.levels[level]
}

func (elem *Element) PrevLevel(level int) *Element {
	if level < 0 || level >= len(elem.levels) {
		return nil
	}
	if level == 0 {
		return elem.prevTopLevel
	}
	prev := elem.prev
	for prev != nil {
		if level < len(prev.levels) {
			return prev
		}
		prev = prev.prevTopLevel
	}
	return prev
}

func (elem *Element) Key() interface{} {
	return elem.key
}

// Score returns the score of this element.
// The score is a hint when comparing elements.
// Skip list keeps all elements sorted by score from smallest to largest.
func (elem *Element) Score() float64 {
	return elem.score
}

// Level returns the level of this elem.
func (elem *Element) Level() int {
	return len(elem.levels)
}

func (elem *Element) reset() {
	elem.list = nil
	elem.prev = nil
	elem.prevTopLevel = nil
	elem.levels = nil
}
