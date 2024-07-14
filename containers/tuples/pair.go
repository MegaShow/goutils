package tuples

// Pair is consisting of two elements.
//
// 两个元素的组合.
type Pair[L, R any] struct {
	Left  L
	Right R
}

// Triple is consisting of three elements.
//
// 三个元素的组合.
type Triple[L, M, R any] struct {
	Left   L
	Middle M
	Right  R
}
