package logic

func And[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return false
			}
		}
		return true
	}
}

func Or[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return true
			}
		}
		return false
	}
}

func Not[T any](predicate func(T) bool) func(T) bool {
	return func(value T) bool {
		return !predicate(value)
	}
}

func NOR[T any](predicates ...func(T) bool) func(T) bool {
	// De Morgan's Law: NOR is equivalent to NOT(OR)
	return Not(Or(predicates...))
}

func NAND[T any](predicates ...func(T) bool) func(T) bool {
	// De Morgan's Law: NAND is equivalent to NOT(AND)
	// It is tragic that, originally, NAND was invented as a universal gate.
	// Probably, the NOT and AND operations I use here are converted to NAND gates in hardware.
	return Not(And(predicates...))
}
