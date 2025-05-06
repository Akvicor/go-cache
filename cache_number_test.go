package cache

import "testing"

func TestIncrementWithInt(t *testing.T) {
	tc := NewNumber[string, int](DefaultExpiration, 0)
	tc.Set("tint", 1, DefaultExpiration)
	err := tc.Increment("tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint")
	if !found {
		t.Error("tint was not found")
	}
	if x != 3 {
		t.Error("tint is not 3:", x)
	}
}

func TestIncrementWithInt8(t *testing.T) {
	tc := NewNumber[string, int8](DefaultExpiration, 0)
	tc.Set("tint8", int8(1), DefaultExpiration)
	err := tc.Increment("tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x != 3 {
		t.Error("tint8 is not 3:", x)
	}
}

func TestIncrementWithInt16(t *testing.T) {
	tc := NewNumber[string, int16](DefaultExpiration, 0)
	tc.Set("tint16", int16(1), DefaultExpiration)
	err := tc.Increment("tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x != 3 {
		t.Error("tint16 is not 3:", x)
	}
}

func TestIncrementWithInt32(t *testing.T) {
	tc := NewNumber[string, int32](DefaultExpiration, 0)
	tc.Set("tint32", int32(1), DefaultExpiration)
	err := tc.Increment("tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x != 3 {
		t.Error("tint32 is not 3:", x)
	}
}

func TestIncrementWithInt64(t *testing.T) {
	tc := NewNumber[string, int64](DefaultExpiration, 0)
	tc.Set("tint64", int64(1), DefaultExpiration)
	err := tc.Increment("tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x != 3 {
		t.Error("tint64 is not 3:", x)
	}
}

func TestIncrementWithUint(t *testing.T) {
	tc := NewNumber[string, uint](DefaultExpiration, 0)
	tc.Set("tuint", uint(1), DefaultExpiration)
	err := tc.Increment("tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x != 3 {
		t.Error("tuint is not 3:", x)
	}
}

func TestIncrementWithUintptr(t *testing.T) {
	tc := NewNumber[string, uintptr](DefaultExpiration, 0)
	tc.Set("tuintptr", uintptr(1), DefaultExpiration)
	err := tc.Increment("tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x != 3 {
		t.Error("tuintptr is not 3:", x)
	}
}

func TestIncrementWithUint8(t *testing.T) {
	tc := NewNumber[string, uint8](DefaultExpiration, 0)
	tc.Set("tuint8", uint8(1), DefaultExpiration)
	err := tc.Increment("tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x != 3 {
		t.Error("tuint8 is not 3:", x)
	}
}

func TestIncrementWithUint16(t *testing.T) {
	tc := NewNumber[string, uint16](DefaultExpiration, 0)
	tc.Set("tuint16", uint16(1), DefaultExpiration)
	err := tc.Increment("tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x != 3 {
		t.Error("tuint16 is not 3:", x)
	}
}

func TestIncrementWithUint32(t *testing.T) {
	tc := NewNumber[string, uint32](DefaultExpiration, 0)
	tc.Set("tuint32", uint32(1), DefaultExpiration)
	err := tc.Increment("tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x != 3 {
		t.Error("tuint32 is not 3:", x)
	}
}

func TestIncrementWithUint64(t *testing.T) {
	tc := NewNumber[string, uint64](DefaultExpiration, 0)
	tc.Set("tuint64", uint64(1), DefaultExpiration)
	err := tc.Increment("tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get("tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x != 3 {
		t.Error("tuint64 is not 3:", x)
	}
}

func TestIncrementWithFloat32(t *testing.T) {
	tc := NewNumber[string, float32](DefaultExpiration, 0)
	tc.Set("float32", float32(1.5), DefaultExpiration)
	err := tc.Increment("float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func TestIncrementWithFloat64(t *testing.T) {
	tc := NewNumber[string, float64](DefaultExpiration, 0)
	tc.Set("float64", float64(1.5), DefaultExpiration)
	err := tc.Increment("float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func TestDecrementWithInt(t *testing.T) {
	tc := NewNumber[string, int](DefaultExpiration, 0)
	tc.Set("int", int(5), DefaultExpiration)
	err := tc.Decrement("int", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int")
	if !found {
		t.Error("int was not found")
	}
	if x != 3 {
		t.Error("int is not 3:", x)
	}
}

func TestDecrementWithInt8(t *testing.T) {
	tc := NewNumber[string, int8](DefaultExpiration, 0)
	tc.Set("int8", int8(5), DefaultExpiration)
	err := tc.Decrement("int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x != 3 {
		t.Error("int8 is not 3:", x)
	}
}

func TestDecrementWithInt16(t *testing.T) {
	tc := NewNumber[string, int16](DefaultExpiration, 0)
	tc.Set("int16", int16(5), DefaultExpiration)
	err := tc.Decrement("int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x != 3 {
		t.Error("int16 is not 3:", x)
	}
}

func TestDecrementWithInt32(t *testing.T) {
	tc := NewNumber[string, int32](DefaultExpiration, 0)
	tc.Set("int32", int32(5), DefaultExpiration)
	err := tc.Decrement("int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x != 3 {
		t.Error("int32 is not 3:", x)
	}
}

func TestDecrementWithInt64(t *testing.T) {
	tc := NewNumber[string, int64](DefaultExpiration, 0)
	tc.Set("int64", int64(5), DefaultExpiration)
	err := tc.Decrement("int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x != 3 {
		t.Error("int64 is not 3:", x)
	}
}

func TestDecrementWithUint(t *testing.T) {
	tc := NewNumber[string, uint](DefaultExpiration, 0)
	tc.Set("uint", uint(5), DefaultExpiration)
	err := tc.Decrement("uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint")
	if !found {
		t.Error("uint was not found")
	}
	if x != 3 {
		t.Error("uint is not 3:", x)
	}
}

func TestDecrementWithUintptr(t *testing.T) {
	tc := NewNumber[string, uintptr](DefaultExpiration, 0)
	tc.Set("uintptr", uintptr(5), DefaultExpiration)
	err := tc.Decrement("uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x != 3 {
		t.Error("uintptr is not 3:", x)
	}
}

func TestDecrementWithUint8(t *testing.T) {
	tc := NewNumber[string, uint8](DefaultExpiration, 0)
	tc.Set("uint8", uint8(5), DefaultExpiration)
	err := tc.Decrement("uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x != 3 {
		t.Error("uint8 is not 3:", x)
	}
}

func TestDecrementWithUint16(t *testing.T) {
	tc := NewNumber[string, uint16](DefaultExpiration, 0)
	tc.Set("uint16", uint16(5), DefaultExpiration)
	err := tc.Decrement("uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x != 3 {
		t.Error("uint16 is not 3:", x)
	}
}

func TestDecrementWithUint32(t *testing.T) {
	tc := NewNumber[string, uint32](DefaultExpiration, 0)
	tc.Set("uint32", uint32(5), DefaultExpiration)
	err := tc.Decrement("uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x != 3 {
		t.Error("uint32 is not 3:", x)
	}
}

func TestDecrementWithUint64(t *testing.T) {
	tc := NewNumber[string, uint64](DefaultExpiration, 0)
	tc.Set("uint64", uint64(5), DefaultExpiration)
	err := tc.Decrement("uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x != 3 {
		t.Error("uint64 is not 3:", x)
	}
}

func TestDecrementWithFloat32(t *testing.T) {
	tc := NewNumber[string, float32](DefaultExpiration, 0)
	tc.Set("float32", float32(5.5), DefaultExpiration)
	err := tc.Decrement("float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x != 3.5 {
		t.Error("float32 is not 3:", x)
	}
}

func TestDecrementWithFloat64(t *testing.T) {
	tc := NewNumber[string, float64](DefaultExpiration, 0)
	tc.Set("float64", float64(5.5), DefaultExpiration)
	err := tc.Decrement("float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get("float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func TestIncrementOverflowInt(t *testing.T) {
	tc := NewNumber[string, int8](DefaultExpiration, 0)
	tc.Set("int8", int8(127), DefaultExpiration)
	err := tc.Increment("int8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get("int8")
	int8 := x
	if int8 != -128 {
		t.Error("int8 did not overflow as expected; value:", int8)
	}

}

func TestIncrementOverflowUint(t *testing.T) {
	tc := NewNumber[string, uint8](DefaultExpiration, 0)
	tc.Set("uint8", uint8(255), DefaultExpiration)
	err := tc.Increment("uint8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get("uint8")
	uint8 := x
	if uint8 != 0 {
		t.Error("uint8 did not overflow as expected; value:", uint8)
	}
}

func TestDecrementUnderflowUint(t *testing.T) {
	tc := NewNumber[string, uint8](DefaultExpiration, 0)
	tc.Set("uint8", uint8(0), DefaultExpiration)
	err := tc.Decrement("uint8", 1)
	if err != nil {
		t.Error("Error decrementing int8:", err)
	}
	x, _ := tc.Get("uint8")
	uint8 := x
	if uint8 != 255 {
		t.Error("uint8 did not underflow as expected; value:", uint8)
	}
}

func BenchmarkIncrementInt(b *testing.B) {
	b.StopTimer()
	tc := NewNumber[string, int](DefaultExpiration, 0)
	tc.Set("foo", 0, DefaultExpiration)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Increment("foo", 1)
	}
}
