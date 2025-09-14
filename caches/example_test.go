package caches

import (
	"context"
	"fmt"
	"time"
)

func Example() {
	ctx := context.Background()
	cache := NewSimpleCache(
		WithCleanDuration[int, string](time.Minute),
		WithExpiration[int, string](time.Hour),
		WithLoaderFunc(func(ctx context.Context, key int) (string, error) {
			return fmt.Sprint(key), nil
		}),
	)
	defer func() { _ = cache.Close() }()

	fmt.Println(cache.Get(ctx, 1))
	fmt.Println(cache.Get(ctx, 2))

	_ = cache.Set(ctx, 1, "11")
	fmt.Println(cache.Get(ctx, 1))
	// Output:
	// 1 <nil>
	// 2 <nil>
	// 11 <nil>
}
