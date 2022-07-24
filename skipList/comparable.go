package skipList

type Comparable interface {
	Compare(lhs, rhs interface{}) int
	CalcScore(key interface{}) float64
}

var (
	_ Comparable = GreaterThanFunc(nil)
	_ Comparable = LessThanFunc(nil)
)

type GreaterThanFunc func(lhs, rhs interface{}) int

type LessThanFunc GreaterThanFunc

func (f GreaterThanFunc) Compare(lhs, rhs interface{}) int {
	return f(lhs, rhs)
}

func (f GreaterThanFunc) CalcScore(key interface{}) float64 {
	return 0
}

func (f LessThanFunc) Compare(lhs, rhs interface{}) int {
	return -f(lhs, rhs)
}

func (f LessThanFunc) CalcScore(key interface{}) float64 {
	return 0
}
func Reverse(comparable Comparable) Comparable {
	return reversedComparable{
		comparable: comparable,
	}

}

type reversedComparable struct {
	comparable Comparable
}

var _ Comparable = reversedComparable{}

func (reversed reversedComparable) Compare(lhs, rhs interface{}) int {
	return -reversed.comparable.Compare(lhs, rhs)
}

func (reversed reversedComparable) CalcScore(key interface{}) float64 {
	return -reversed.comparable.CalcScore(key)
}
