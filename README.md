# go-cache

Rewrite https://github.com/patrickmn/go-cache

Usage is similar to go-cache

```go
// New Any
New[string, any](5*time.Minute, 0)
NewAny[string, any](5*time.Minute, 0)
// New Number
NewNumber[string, any](5*time.Minute, 0)
```

- Any: `[K comparable, V any]` Allows any type as a value
- Number: `[K comparable, V number]` Allows numeric types as values

```
type number interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 |
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
  ~uintptr | ~float32 | ~float64
}
```

