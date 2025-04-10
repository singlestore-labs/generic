package generic_test

// This file generated with Claude 3.7 Sonnet

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/singlestore-labs/generic"
)

func TestCopySlice(t *testing.T) {
	t.Parallel()

	t.Run("copies int slice", func(t *testing.T) {
		t.Parallel()

		original := []int{1, 2, 3, 4, 5}
		copied := generic.CopySlice(original)

		t.Log("Copied slice should equal original but be a distinct slice")
		assert.Equal(t, original, copied)

		t.Log("Modifying original should not affect copy")
		original[0] = 99
		assert.NotEqual(t, original[0], copied[0])
	})

	t.Run("handles empty slice", func(t *testing.T) {
		t.Parallel()

		original := []string{}
		copied := generic.CopySlice(original)

		t.Log("Should return empty slice when copying empty slice")
		assert.Empty(t, copied)
		assert.Len(t, copied, 0)
	})

	t.Run("copies struct slice", func(t *testing.T) {
		t.Parallel()

		type person struct {
			Name string
			Age  int
		}

		original := []person{
			{"Alice", 30},
			{"Bob", 25},
		}
		copied := generic.CopySlice(original)

		t.Log("Should copy structs correctly")
		assert.Equal(t, original, copied)
	})
}

func TestCastStringySlice(t *testing.T) {
	t.Parallel()

	t.Run("cast string to string", func(t *testing.T) {
		t.Parallel()

		original := []string{"hello", "world"}
		result := generic.CastStringySlice[string](original)

		t.Log("Should maintain same values when casting to same type")
		assert.Equal(t, original, result)
	})

	t.Run("cast rune slice to string", func(t *testing.T) {
		t.Parallel()

		original := [][]rune{[]rune("hello"), []rune("world")}
		result := generic.CastStringySlice[string](original)

		t.Log("Should convert rune slices to strings")
		expected := []string{"hello", "world"}
		assert.Equal(t, expected, result)
	})

	t.Run("cast string to rune slice", func(t *testing.T) {
		t.Parallel()

		original := []string{"abc", "def"}
		result := generic.CastStringySlice[[]rune](original)

		t.Log("Should convert strings to rune slices")
		expected := [][]rune{[]rune("abc"), []rune("def")}
		assert.Equal(t, expected, result)
	})
}

func TestTransformSlice(t *testing.T) {
	t.Parallel()

	t.Run("transforms integers to strings", func(t *testing.T) {
		t.Parallel()

		original := []int{1, 2, 3}
		result := generic.TransformSlice(original, strconv.Itoa)

		t.Log("Should convert ints to their string representation")
		expected := []string{"1", "2", "3"}
		assert.Equal(t, expected, result)
	})

	t.Run("transforms with custom function", func(t *testing.T) {
		t.Parallel()

		original := []int{1, 2, 3}
		result := generic.TransformSlice(original, func(i int) int {
			return i * i
		})

		t.Log("Should apply squaring function to each element")
		expected := []int{1, 4, 9}
		assert.Equal(t, expected, result)
	})

	t.Run("transforms with struct conversion", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
		}

		type Employee struct {
			Person
			ID int
		}

		original := []Person{
			{"Alice"},
			{"Bob"},
		}

		result := generic.TransformSlice(original, func(p Person) Employee {
			return Employee{Person: p, ID: len(p.Name)}
		})

		t.Log("Should transform Person to Employee")
		expected := []Employee{
			{Person{"Alice"}, 5},
			{Person{"Bob"}, 3},
		}
		assert.Equal(t, expected, result)
	})
}

func TestSliceContains(t *testing.T) {
	t.Parallel()

	t.Run("finds matching element", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5}
		result := generic.SliceContains(slice, func(i int) bool {
			return i == 3
		})

		t.Log("Should return true when slice contains matching element")
		assert.True(t, result)
	})

	t.Run("returns false for no match", func(t *testing.T) {
		t.Parallel()

		slice := []string{"apple", "banana", "cherry"}
		result := generic.SliceContains(slice, func(s string) bool {
			return s == "grape"
		})

		t.Log("Should return false when no element matches")
		assert.False(t, result)
	})

	t.Run("checks complex condition", func(t *testing.T) {
		t.Parallel()

		type User struct {
			Name  string
			Admin bool
		}

		users := []User{
			{"Alice", false},
			{"Bob", true},
			{"Charlie", false},
		}

		result := generic.SliceContains(users, func(u User) bool {
			return u.Admin
		})

		t.Log("Should return true when slice contains an admin user")
		assert.True(t, result)
	})
}

func TestSliceEvery(t *testing.T) {
	t.Parallel()

	t.Run("all elements match", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 1, 1, 1}
		returns := generic.AllElements(slice, func(i int) bool {
			return i == 1
		})

		t.Log("Should return true when all elements match condition")
		assert.True(t, returns)
	})

	t.Run("returns false if not all elements match", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 1, 1}
		returns := generic.AllElements(slice, func(i int) bool {
			return i == 1
		})

		t.Log("Should return false when not all elements match condition")
		assert.False(t, returns)
	})

	t.Run("checks complex condition", func(t *testing.T) {
		t.Parallel()

		type User struct {
			Name  string
			Admin bool
		}

		users := []User{
			{"Alice", false},
			{"Bob", false},
			{"Charlie", false},
		}

		returns := generic.AllElements(users, func(u User) bool {
			return !u.Admin
		})

		t.Log("Should return true when all users are not admins")
		assert.True(t, returns)

		users = []User{
			{"Alice", false},
			{"Bob", true},
			{"Charlie", false},
		}

		returns = generic.AllElements(users, func(u User) bool {
			return u.Admin
		})

		t.Log("Should return false when not all users are admins")
		assert.False(t, returns)
	})
}

func TestSliceContainsElement(t *testing.T) {
	t.Parallel()

	t.Run("finds int element", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5}

		t.Log("Should find existing element")
		assert.True(t, generic.SliceContainsElement(slice, 3))

		t.Log("Should not find non-existing element")
		assert.False(t, generic.SliceContainsElement(slice, 6))
	})

	t.Run("finds string element", func(t *testing.T) {
		t.Parallel()

		slice := []string{"apple", "banana", "cherry"}

		t.Log("Should find existing element")
		assert.True(t, generic.SliceContainsElement(slice, "banana"))

		t.Log("Should not find non-existing element")
		assert.False(t, generic.SliceContainsElement(slice, "grape"))
	})

	t.Run("handles empty slice", func(t *testing.T) {
		t.Parallel()

		slice := []int{}

		t.Log("Should return false for empty slice")
		assert.False(t, generic.SliceContainsElement(slice, 1))
	})
}

func TestCountMatchingElements(t *testing.T) {
	t.Parallel()

	t.Run("counts even numbers", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5, 6}
		count := generic.CountMatchingElements(slice, func(i int) bool {
			return i%2 == 0
		})

		t.Log("Should count 3 even numbers in the slice")
		assert.Equal(t, 3, count)
	})

	t.Run("counts no matches", func(t *testing.T) {
		t.Parallel()

		slice := []string{"apple", "banana", "cherry"}
		count := generic.CountMatchingElements(slice, func(s string) bool {
			return len(s) > 10
		})

		t.Log("Should count 0 strings longer than 10 characters")
		assert.Equal(t, 0, count)
	})

	t.Run("handles empty slice", func(t *testing.T) {
		t.Parallel()

		slice := []int{}
		count := generic.CountMatchingElements(slice, func(i int) bool {
			return true
		})

		t.Log("Should return 0 for empty slice")
		assert.Equal(t, 0, count)
	})
}

func TestCombineSlices(t *testing.T) {
	t.Parallel()

	t.Run("combines multiple slices", func(t *testing.T) {
		t.Parallel()

		s1 := []int{1, 2}
		s2 := []int{3, 4}
		s3 := []int{5, 6}

		result := generic.CombineSlices(s1, s2, s3)

		t.Log("Should combine all slices in order")
		expected := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, expected, result)
	})

	t.Run("returns first slice if only one provided", func(t *testing.T) {
		t.Parallel()

		s1 := []string{"a", "b", "c"}
		result := generic.CombineSlices(s1)

		t.Log("Should return first slice when only one is provided")
		assert.Equal(t, s1, result)
	})

	t.Run("handles empty first slice", func(t *testing.T) {
		t.Parallel()

		s1 := []int{}
		s2 := []int{1, 2, 3}

		result := generic.CombineSlices(s1, s2)

		t.Log("Should return second slice when first is empty")
		assert.Equal(t, s2, result)
	})

	t.Run("handles all empty slices", func(t *testing.T) {
		t.Parallel()

		s1 := []int{}
		s2 := []int{}

		result := generic.CombineSlices(s1, s2)

		t.Log("Should return empty slice when all inputs are empty")
		assert.Empty(t, result)
	})
}

func TestPrepend(t *testing.T) {
	t.Parallel()

	t.Run("prepends elements to slice", func(t *testing.T) {
		t.Parallel()

		existing := []int{3, 4, 5}
		newHead := []int{1, 2}

		result := generic.Prepend(existing, newHead...)

		t.Log("Should prepend elements to the beginning of slice")
		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("prepends to empty slice", func(t *testing.T) {
		t.Parallel()

		existing := []string{}
		newHead := []string{"a", "b"}

		result := generic.Prepend(existing, newHead...)

		t.Log("Should return newHead when existing is empty")
		assert.Equal(t, newHead, result)
	})

	t.Run("handles empty prepend", func(t *testing.T) {
		t.Parallel()

		existing := []int{1, 2, 3}

		result := generic.Prepend(existing)

		t.Log("Should return existing when nothing to prepend")
		assert.Equal(t, existing, result)
	})
}

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	t.Run("removes duplicate integers", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 2, 3, 4, 4, 5}
		result := generic.RemoveDuplicates(slice)

		t.Log("Should remove duplicate integers while preserving order")
		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("removes duplicate strings", func(t *testing.T) {
		t.Parallel()

		slice := []string{"a", "b", "b", "c", "a", "d"}
		result := generic.RemoveDuplicates(slice)

		t.Log("Should remove duplicate strings while preserving order")
		expected := []string{"a", "b", "c", "d"}
		assert.Equal(t, expected, result)
	})

	t.Run("handles no duplicates", func(t *testing.T) {
		t.Parallel()

		slice := []int{5, 4, 3, 2, 1}
		result := generic.RemoveDuplicates(slice)

		t.Log("Should return same elements when no duplicates exist")
		expected := []int{5, 4, 3, 2, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("handles empty slice", func(t *testing.T) {
		t.Parallel()

		slice := []int{}
		result := generic.RemoveDuplicates(slice)

		t.Log("Should return empty slice for empty input")
		assert.Empty(t, result)
	})
}

func TestDeleteFromSlice(t *testing.T) {
	t.Parallel()

	t.Run("deletes middle element", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5}
		result := generic.DeleteFromSlice(slice, 2)

		t.Log("Should delete element at index 2 (value 3)")
		t.Log("Note: This function reorders the original slice")
		assert.Len(t, result, 4)
		assert.NotContains(t, result, 3)
	})

	t.Run("deletes first element", func(t *testing.T) {
		t.Parallel()

		slice := []string{"a", "b", "c"}
		result := generic.DeleteFromSlice(slice, 0)

		t.Log("Should delete element at index 0 (value 'a')")
		assert.Len(t, result, 2)
		assert.NotContains(t, result, "a")
	})

	t.Run("deletes last element", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3}
		result := generic.DeleteFromSlice(slice, 2)

		t.Log("Should delete element at index 2 (value 3)")
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})
}

func TestIntersectSlices(t *testing.T) {
	t.Parallel()

	t.Run("finds common elements", func(t *testing.T) {
		t.Parallel()

		a := []int{1, 2, 3, 4, 5}
		b := []int{3, 4, 5, 6, 7}

		result := generic.IntersectSlices(a, b)

		t.Log("Should return elements common to both slices in order of b")
		expected := []int{3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("handles no common elements", func(t *testing.T) {
		t.Parallel()

		a := []string{"a", "b", "c"}
		b := []string{"d", "e", "f"}

		result := generic.IntersectSlices(a, b)

		t.Log("Should return empty slice when no common elements")
		assert.Empty(t, result)
	})

	t.Run("handles empty slices", func(t *testing.T) {
		t.Parallel()

		a := []int{}
		b := []int{1, 2, 3}

		result := generic.IntersectSlices(a, b)

		t.Log("Should return empty slice when one input is empty")
		assert.Empty(t, result)
	})

	t.Run("preserves order of second slice", func(t *testing.T) {
		t.Parallel()

		a := []int{5, 3, 1, 2, 4}
		b := []int{4, 2, 3, 1}

		result := generic.IntersectSlices(a, b)

		t.Log("Should preserve order of b slice in result")
		expected := []int{4, 2, 3, 1}
		assert.Equal(t, expected, result)
	})
}
