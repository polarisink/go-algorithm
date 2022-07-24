package skipList

import "math/rand"

var DefaultMaxLevel = 48

const preallocDefaultMaxLevel = 48

type SkipList struct {
	elementHeader

	comparable Comparable
	rand       *rand.Rand

	maxLevel int
	length   int
	back     *Element
}

func New(comparable Comparable) *SkipList {
	if DefaultMaxLevel <= 0 {
		panic("skiplist default level must not be zero or negative")
	}
	return nil
}
