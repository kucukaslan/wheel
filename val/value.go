package val

// ToPtr returns a pointer to the value v.
func ToPtr[T any](v *T) T {
	return *v
}
