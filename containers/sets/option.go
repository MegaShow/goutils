package sets

// Option is option for set.
//
// 集合选项.
type Option struct {
	capacity int
}

// WithCapacity sets a initial capacity for set.
//
// 为集合设置一个初始化容量.
func WithCapacity(capacity int) Option {
	return Option{
		capacity: capacity,
	}
}
