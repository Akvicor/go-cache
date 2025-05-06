package cache

import "time"

type Item[V any] struct {
	Value      V
	Hit        int   // 命中次数
	Expiration int64 // 过期时间
}

// IsHit 判断是否命中过
func (item Item[V]) IsHit() bool {
	return item.Hit > 0
}

// Expired 判断是否过期
func (item Item[V]) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}
