package generic

// Keys returns the map keys as a slice
func Keys[K comparable, V any](m map[K]V) []K {
	slice := make([]K, 0, len(m))
	for k := range m {
		slice = append(slice, k)
	}
	return slice
}

// Values returns the map values as a slice
func Values[K comparable, V any](m map[K]V) []V {
	slice := make([]V, 0, len(m))
	for _, v := range m {
		slice = append(slice, v)
	}
	return slice
}

// CompareKeys returns slice of keys that are only in a
// and a slice of keys that are only in b.
func CompareKeys[K comparable, V any](a, b map[K]V) ([]K, []K) {
	return MissingKeys(a, b), MissingKeys(b, a)
}

// MissingKeys returns the keys that are in a but not b
func MissingKeys[K comparable, V any](a, b map[K]V) []K {
	// Pre-allocate with capacity of a since that's the maximum possible size
	onlyA := make([]K, 0, len(a))
	for k := range a {
		if _, ok := b[k]; !ok {
			onlyA = append(onlyA, k)
		}
	}
	return onlyA
}

// EqualKeys checks if two maps have exactly the same keys.
// Returns true if both maps contain the same set of keys, regardless of values.
func EqualKeys[K comparable, V any](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	// We don't need to check if there are extras in b not in a because 
	// we checked that there are an equal number of keys so all we have
	// to check is that all of a is in b
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	// Pre-allocate with capacity of m
	newM := make(map[K]V, len(m))
	for k, v := range m {
		newM[k] = v
	}
	return newM
}

// CopyMapSubset creates a new map containing only the specified keys from the original map
func CopyMapSubset[K comparable, V any](m map[K]V, keys []K) map[K]V {
	if m == nil {
		return nil
	}

	if len(m) == 0 {
		return make(map[K]V)
	}

	// Pre-allocate with capacity of keys since that's the maximum possible size
	result := make(map[K]V, len(keys))
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Merge copies b onto a, overriding any common keys:
// a is modified and returned.
func Merge[K comparable, V any](a, b map[K]V) map[K]V {
	if a == nil {
		return CopyMap(b)
	}
	for k, v := range b {
		a[k] = v
	}
	return a
}

// AllKeys returns true if all keys in the map satisfy the given filter function.
// Returns true for empty maps (vacuous truth).
func AllKeys[K comparable, V any](m map[K]V, filter func(K) bool) bool {
	for k := range m {
		if !filter(k) {
			return false
		}
	}
	return true
}

// AnyKey returns true if at least one key in the map satisfies the given filter function.
// Returns false for empty maps.
func AnyKey[K comparable, V any](m map[K]V, filter func(K) bool) bool {
	for k := range m {
		if filter(k) {
			return true
		}
	}
	return false
}

// AllValues returns true if all values in the map satisfy the given filter function.
// Returns true for empty maps (vacuous truth).
func AllValues[K comparable, V any](m map[K]V, filter func(V) bool) bool {
	for _, v := range m {
		if !filter(v) {
			return false
		}
	}
	return true
}

// AnyValue returns true if at least one value in the map satisfies the given filter function.
// Returns false for empty maps.
func AnyValue[K comparable, V any](m map[K]V, filter func(V) bool) bool {
	for _, v := range m {
		if filter(v) {
			return true
		}
	}
	return false
}
