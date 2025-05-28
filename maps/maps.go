package maps

import (
	"maps"
	"slices"
)

// Keys returns the keys of the map as a slice.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// KeysButCooler returns the keys of the map just like the Keys.
// However, instead of using a for loop, it uses the `maps.Keys` to obtain an iter.Seq[K]
// and then collects the keys into a slice using `slices.Collect`.
// It's probably worse due to indirections, but it is just seems cooler to do it this way.
func KeysButCooler[K comparable, V any](m map[K]V) []K {
	return slices.Collect(maps.Keys(m))
}

// Values returns the values of the map as a slice.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ValuesButCooler returns the values of the map just like the Values.
// However, instead of using a for loop, it uses the `maps.Values` to obtain an iter.Seq[V]
// and then collects the values into a slice using `slices.Collect`.
func ValuesButCooler[K comparable, V any](m map[K]V) []V {
	return slices.Collect(maps.Values(m))
}

func UniqueValues[K comparable, V comparable](m map[K]V) []V {
	// Use maps.Values to get unique values
	values := make(map[V]bool, len(m))
	for _, v := range m {
		values[v] = true
	}
	return Keys(values)
}

func ToSlice[K comparable, V any](m map[K]V) []struct {
	Key   K
	Value V
} {
	slice := make([]struct {
		Key   K
		Value V
	}, 0, len(m))
	for k, v := range m {
		slice = append(slice, struct {
			Key   K
			Value V
		}{Key: k, Value: v})
	}
	return slice
}
