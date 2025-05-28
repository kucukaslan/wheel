package ptr

// ToVal returns the value pointed to by v.
func ToVal[T any](v *T) T {
	return *v
}
