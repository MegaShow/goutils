package internal

import "encoding/json"

// Container is the universal interface for containers.
//
// 容器通用接口.
type Container[T any] interface {
	Clear()
	Len() int
	Values() []T

	json.Marshaler
	json.Unmarshaler
}
