package utils

// NewBool creates a new bool pointer from a bool value.
func NewBool(b bool) *bool {
	v := b
	return &v
}

// NewInt creates a new int pointer from an int value.
func NewInt(i int) *int {
	v := i
	return &v
}

// NewString creates a new string pointer from a string value.
func NewString(s string) *string {
	v := s
	return &v
}

// NewFloat64 creates a new float64 pointer from a float64 value.
func NewFloat64(f float64) *float64 {
	v := f
	return &v
}
