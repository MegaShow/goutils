package sets

import "encoding/json"

var _ Set[int] = (*HashSet[int])(nil)

// HashSet is a unordered set implemented by map,
//
// 无序的哈希集合, 由 map 实现.
type HashSet[T comparable] struct {
	m map[T]struct{}
}

// NewHashSet create a hash set.
//
// 创建一个哈希集合.
func NewHashSet[T comparable](opts ...Option) *HashSet[T] {
	s := new(HashSet[T])

	for _, opt := range opts {
		if opt.capacity != 0 {
			s.m = make(map[T]struct{}, opt.capacity)
		}
	}

	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	return s
}

// Add inserts an item.
//
// 插入新元素.
func (s *HashSet[T]) Add(item T) {
	s.m[item] = struct{}{}
}

// Contains returns whether set contains this item.
//
// 判断集合是否包含该元素.
func (s *HashSet[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

// Remove removes item from set, returns false if set does not contain thie item.
//
// 从集合中移除元素, 如果集合中不包含该元素将返回 false.
func (s *HashSet[T]) Remove(item T) bool {
	_, ok := s.m[item]
	if ok {
		delete(s.m, item)
	}
	return ok
}

// Clear removes all items of set.
//
// 移除集合的所有元素.
func (s *HashSet[T]) Clear() {
	clear(s.m)
}

// Len returns length of set.
//
// 返回集合元素的数量.
func (s *HashSet[T]) Len() int {
	return len(s.m)
}

// Values returns all items as slice.
//
// 返回包含所有元素的切片.
func (s *HashSet[T]) Values() []T {
	items := make([]T, 0, len(s.m))
	for item := range s.m {
		items = append(items, item)
	}
	return items
}

// MarshalJSON marshal set into json bytes.
//
// 序列化集合.
func (s *HashSet[T]) MarshalJSON() ([]byte, error) {
	values := s.Values()
	return json.Marshal(values)
}

// UnmarshalJSON unmarshal json bytes to set.
//
// 反序列化集合.
func (s *HashSet[T]) UnmarshalJSON(data []byte) error {
	values := make([]T, 0)
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}
	if s.m == nil {
		s.m = make(map[T]struct{}, len(values))
	}
	for _, value := range values {
		s.m[value] = struct{}{}
	}
	return nil
}
