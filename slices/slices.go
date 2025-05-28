package slices

import (
	"slices"
	"wheel-go/logic"
)

// FilterSlice gives a new slice containing only the elements of the input slice
// that satisfy all provided filters.
func FilterSlice[T any](slice []T, filters ...func(T) bool) []T {
	if slice == nil {
		return nil
	}
	var result = make([]T, 0, len(slice))
to_next_item:
	for _, item := range slice {
		for _, filter := range filters {
			// If any filter returns false, skip this item
			if !filter(item) {
				continue to_next_item
			}
		}
		result = append(result, item)
	}
	return result
}

// FilterSliceInplace filters a slice in place, modifying the original slice.
// Unlike the slices.DeleteFunc, this function accepts multiple filters.
// Technically, it is possible to write a wrapper around slices.DeleteFunc to function
// using logic.And. See FilterSliceInplaceWithDeleteFunc
func FilterSliceInplace[T any](slice []T, filters ...func(T) bool) []T {
	if slice == nil {
		return nil
	}
	var j int
	for i := 0; i < len(slice); i++ {
		item := slice[i]
		include := true
		for _, filter := range filters {
			if !filter(item) {
				include = false
				break
			}
		}
		if include {
			slice[j] = item
			j++
		}
	}
	clear(slice[j:]) // Clear the remaining elements
	return slice[:j]
}

func FilterSliceInplaceWithDeleteFunc[T any](slice []T, filters ...func(T) bool) []T {
	return slices.DeleteFunc(slice, logic.And(filters...))
}
