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
	onlyA := make([]K, 0, len(a))
	for k := range a {
		if _, ok := b[k]; !ok {
			onlyA = append(onlyA, k)
		}
	}
	return onlyA
}

func EqualKeys[K comparable, V any](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	if len(MissingKeys(a, b)) != 0 {
		return false
	}
	return true
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	newM := make(map[K]V)
	for k, v := range m {
		newM[k] = v
	}
	return newM
}

// Merge copies b onto a, overriding any common keys:
// a is modified and returned.
func Merge[K comparable, V any](a, b map[K]V) map[K]V {
	if a == nil {
		return b
	}
	for k, v := range b {
		a[k] = v
	}
	return a
}

func AllKeys[K comparable, V any](m map[K]V, filter func(K) bool) bool {
	for k := range m {
		if !filter(k) {
			return false
		}
	}
	return true
}

func AnyKey[K comparable, V any](m map[K]V, filter func(K) bool) bool {
	for k := range m {
		if filter(k) {
			return true
		}
	}
	return false
}

func AllValues[K comparable, V any](m map[K]V, filter func(V) bool) bool {
	for _, v := range m {
		if !filter(v) {
			return false
		}
	}
	return true
}

func AnyValue[K comparable, V any](m map[K]V, filter func(V) bool) bool {
	for _, v := range m {
		if filter(v) {
			return true
		}
	}
	return false
}
