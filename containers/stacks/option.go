package stacks

// Option is option for stack.
//
// 栈选项.
type Option struct {
	capacity int
}

// WithCapacity sets a initial capacity for stack.
//
// 为栈设置一个初始化容量.
func WithCapacity(capacity int) Option {
	return Option{
		capacity: capacity,
	}
}
