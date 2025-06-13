package funcs

func Identity(element any) any {
	return element
}

func IdentityTypeSafe[T any](element T) T {
	return element
}

func True(element any) bool {
	return true
}

func TrueVariadic(elements ...any) bool {
	return true
}

func False(element any) bool {
	return false
}

func FalseVariadic(elements ...any) bool {
	return false
}

func ZeroValue[T any](element T) T {
	var zero T
	return zero
}
