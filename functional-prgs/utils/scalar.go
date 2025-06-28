package utils

// Scalar is a type constraint that allows any numeric type. The ~ operator allows types that have the specified
// underlying type (e.g., custom types based on int). The | operator is used to combine multiple types. This constraint
// enables generic functions to accept any integer or floating-point type, including user-defined types with those
// underlying types.
type Scalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func AddScalar[T Scalar](x, y T) T {
	return x + y
}
