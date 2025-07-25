package generic

func CopySlice[T any](orig []T) []T {
	c := make([]T, len(orig))
	copy(c, orig)
	return c
}

// FilterSlice returns a new slice containing only elements for which filter returns true.
func FilterSlice[T any](slice []T, filter func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if filter(item) {
			result = append(result, item)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func CastStringySlice[B, A ~string | ~[]rune](orig []A) []B {
	c := make([]B, len(orig))
	for i, a := range orig {
		c[i] = B(a)
	}
	return c
}

func TransformSlice[T any, U any](orig []T, cast func(T) U) []U {
	c := make([]U, len(orig))
	for i, a := range orig {
		c[i] = cast(a)
	}
	return c
}

func SliceContains[T any](slice []T, filter func(t T) bool) bool {
	for _, item := range slice {
		if filter(item) {
			return true
		}
	}
	return false
}

func AllElements[T any](slice []T, filter func(t T) bool) bool {
	for _, item := range slice {
		if !filter(item) {
			return false
		}
	}
	return true
}

func SliceContainsElement[T comparable](slice []T, element T) bool {
	return SliceContains(slice, func(t T) bool {
		return t == element
	})
}

func CountMatchingElements[T any](s []T, filter func(T) bool) int {
	var c int
	for _, e := range s {
		if filter(e) {
			c++
		}
	}
	return c
}

// FirstMatchIndex returns -1 if there are no matches
func FirstMatchIndex[T any](s []T, filter func(T) bool) int {
	for i, e := range s {
		if filter(e) {
			return i
		}
	}
	return -1
}

// ReplaceFirstMatchOrAppend appends the new element unless there is a matching element
func ReplaceOrAppend[T any](s []T, n T, filter func(T) bool) []T {
	i := FirstMatchIndex(s, filter)
	c := CopySlice(s)
	if i != -1 {
		c[i] = n
	} else {
		c = append(c, n)
	}
	return c
}

// CombineSlices may return the first slice if it is the only slice with elements.  A copy
// is only made if it has to be made. For no input, nil is returned.
func CombineSlices[T any](slices ...[]T) []T {
	switch len(slices) {
	case 0:
		return nil
	case 1:
		return slices[0]
	}
	return CombineSlicesCopy[T](slices...)
}

// CombineSlicesCopy combines all of its input slices, always
// returning a new slice, For no input, an empty slice is returned.
func CombineSlicesCopy[T any](slices ...[]T) []T {
	var total int
	for _, m := range slices {
		total += len(m)
	}
	combined := make([]T, 0, total)
	for _, m := range slices {
		combined = append(combined, m...)
	}
	return combined
}

func Prepend[T any](existing []T, newHead ...T) []T {
	return CombineSlices(newHead, existing)
}

// RemoveDuplicates removes duplicate members of a slice. Slice order is preserved.
// RemoveDuplicates works only on slices of comparable types. It does not support
// slices of interfaces.
func RemoveDuplicates[T comparable](existing []T) []T {
	if len(existing) <= 1 {
		return existing
	}
	new := make([]T, 0, len(existing))
	values := make(map[T]bool)
	for _, value := range existing {
		if !values[value] {
			new = append(new, value)
			values[value] = true
		}
	}
	return new
}

// DeleteFromSlice removes an item from a slice. It may reorder the original slice
// in the process. It executes in O(1) time.
func DeleteFromSlice[T comparable](prior []T, index int) []T {
	prior[index] = prior[len(prior)-1]
	return prior[:len(prior)-1]
}

// IntersectSlices returns elements in common to two slices ordered as per the
// order of b
func IntersectSlices[T comparable](a []T, b []T) []T {
	m := ToSet(a)
	u := make([]T, 0, len(b))
	for _, e := range b {
		if _, ok := m[e]; ok {
			u = append(u, e)
		}
	}
	return u
}
