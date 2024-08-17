package caches

import (
	"context"
	"time"
)

type baseCache[K comparable, V any] struct {
	cleanDuration time.Duration
	expiration    time.Duration
	loaderFunc    func(ctx context.Context, key K) (V, error)
}

func (c *baseCache[K, V]) initOpts(opts []Option[K, V]) {
	c.cleanDuration = time.Minute // 默认1分钟
	for _, opt := range opts {
		opt(c)
	}
}
