package generic_test

// This file generated with Claude 3.7 Sonnet

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/singlestore-labs/generic"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	t.Run("extracts keys from map", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		keys := generic.Keys(m)

		t.Log("Should extract all keys from map")
		assert.Len(t, keys, 3)
		assert.Contains(t, keys, "a")
		assert.Contains(t, keys, "b")
		assert.Contains(t, keys, "c")
	})

	t.Run("handles empty map", func(t *testing.T) {
		t.Parallel()

		m := map[int]string{}
		keys := generic.Keys(m)

		t.Log("Should return empty slice for empty map")
		assert.Empty(t, keys)
	})
}

func TestCompareKeys(t *testing.T) {
	t.Parallel()

	t.Run("finds keys in each map", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		b := map[string]int{
			"b": 20,
			"c": 30,
			"d": 40,
		}

		onlyA, onlyB := generic.CompareKeys(a, b)

		t.Log("Should find keys unique to each map")
		assert.Len(t, onlyA, 1)
		assert.Contains(t, onlyA, "a")

		assert.Len(t, onlyB, 1)
		assert.Contains(t, onlyB, "d")
	})

	t.Run("handles identical maps", func(t *testing.T) {
		t.Parallel()

		a := map[int]string{
			1: "a",
			2: "b",
		}

		b := map[int]string{
			1: "x",
			2: "y",
		}

		onlyA, onlyB := generic.CompareKeys(a, b)

		t.Log("Should return empty slices when maps have same keys")
		assert.Empty(t, onlyA)
		assert.Empty(t, onlyB)
	})

	t.Run("handles empty maps", func(t *testing.T) {
		t.Parallel()

		a := map[string]bool{}

		b := map[string]bool{
			"x": true,
			"y": false,
		}

		onlyA, onlyB := generic.CompareKeys(a, b)

		t.Log("Should handle when one map is empty")
		assert.Empty(t, onlyA)
		assert.Len(t, onlyB, 2)
		assert.Contains(t, onlyB, "x")
		assert.Contains(t, onlyB, "y")
	})
}

func TestMissingKeys(t *testing.T) {
	t.Parallel()

	t.Run("finds keys in a but not b", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		b := map[string]int{
			"b": 20,
			"c": 30,
		}

		missing := generic.MissingKeys(a, b)

		t.Log("Should find keys in a that are not in b")
		assert.Len(t, missing, 1)
		assert.Contains(t, missing, "a")
	})

	t.Run("returns empty when no missing keys", func(t *testing.T) {
		t.Parallel()

		a := map[int]string{
			1: "a",
			2: "b",
		}

		b := map[int]string{
			1: "x",
			2: "y",
			3: "z",
		}

		missing := generic.MissingKeys(a, b)

		t.Log("Should return empty slice when all keys in a are also in b")
		assert.Empty(t, missing)
	})

	t.Run("handles empty maps", func(t *testing.T) {
		t.Parallel()

		a := map[string]bool{}
		b := map[string]bool{"x": true}

		missing := generic.MissingKeys(a, b)

		t.Log("Should return empty slice when a is empty")
		assert.Empty(t, missing)

		a = map[string]bool{"y": false}
		b = map[string]bool{}

		missing = generic.MissingKeys(a, b)

		t.Log("Should return all keys from a when b is empty")
		assert.Len(t, missing, 1)
		assert.Contains(t, missing, "y")
	})
}

func TestEqualKeys(t *testing.T) {
	t.Parallel()

	t.Run("equal keys returns true", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
		}

		b := map[string]int{
			"a": 10,
			"b": 20,
		}

		equal := generic.EqualKeys(a, b)

		t.Log("Should return true when maps have same keys")
		assert.True(t, equal)
	})

	t.Run("different keys returns false", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
		}

		b := map[string]int{
			"b": 20,
			"c": 30,
		}

		equal := generic.EqualKeys(a, b)

		t.Log("Should return false when maps have different keys")
		assert.False(t, equal)
	})

	t.Run("different lengths returns false", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
		}

		b := map[string]int{
			"a": 10,
			"b": 20,
			"c": 30,
		}

		equal := generic.EqualKeys(a, b)

		t.Log("Should return false when maps have different number of keys")
		assert.False(t, equal)
	})

	t.Run("empty maps are equal", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{}
		b := map[string]int{}

		equal := generic.EqualKeys(a, b)

		t.Log("Should return true when both maps are empty")
		assert.True(t, equal)
	})
}

func TestCopyMap(t *testing.T) {
	t.Parallel()

	t.Run("copies map contents", func(t *testing.T) {
		t.Parallel()

		original := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		copied := generic.CopyMap(original)

		t.Log("Copied map should equal original but be distinct")
		assert.Equal(t, original, copied)

		t.Log("Modifying original should not affect copy")
		original["a"] = 100
		assert.NotEqual(t, original["a"], copied["a"])
	})

	t.Run("returns nil for nil input", func(t *testing.T) {
		t.Parallel()

		var original map[string]int = nil
		copied := generic.CopyMap(original)

		t.Log("Should return nil when input map is nil")
		assert.Nil(t, copied)
	})

	t.Run("handles empty map", func(t *testing.T) {
		t.Parallel()

		original := map[int]string{}
		copied := generic.CopyMap(original)

		t.Log("Should return empty map for empty input")
		assert.Empty(t, copied)
		assert.NotNil(t, copied)
	})
}

func TestMerge(t *testing.T) {
	t.Parallel()

	t.Run("merges maps", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
		}

		b := map[string]int{
			"b": 20,
			"c": 30,
		}

		result := generic.Merge(a, b)

		t.Log("Should merge b into a with b's values taking precedence")
		expected := map[string]int{
			"a": 1,
			"b": 20,
			"c": 30,
		}
		assert.Equal(t, expected, result)
		assert.Equal(t, a, result) // Merge returns a
	})

	t.Run("handles nil first map", func(t *testing.T) {
		t.Parallel()

		var a map[string]int = nil
		b := map[string]int{
			"x": 10,
			"y": 20,
		}

		result := generic.Merge(a, b)

		t.Log("Should return b when a is nil")
		assert.Equal(t, b, result)
	})

	t.Run("handles empty second map", func(t *testing.T) {
		t.Parallel()

		a := map[string]int{
			"a": 1,
			"b": 2,
		}

		b := map[string]int{}

		result := generic.Merge(a, b)

		t.Log("Should return a unchanged when b is empty")
		assert.Equal(t, a, result)
	})
}

func TestAllKeys(t *testing.T) {
	t.Parallel()

	t.Run("returns true if all keys satisfy a condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"test_a": 1,
			"test_b": 2,
			"test_c": 3,
		}

		result := generic.AllKeys(m, func(k string) bool {
			return strings.HasPrefix(k, "test_");
		})

		t.Log("Should return true if all keys satisfy the condition")
		assert.True(t, result)
	})

	t.Run("returns false if any key does not satisfy the condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"test_a": 1,
			"test_b": 2,
			"other_c": 3,
		}

		result := generic.AllKeys(m, func(k string) bool {
			return strings.HasPrefix(k, "test_");
		})

		t.Log("Should return false if any key does not satisfy the condition")
		assert.False(t, result)
	})
}

func TestAnyKey(t *testing.T) {
	t.Parallel()

	t.Run("returns true if any key satisfies a condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": 1,
			"test_b": 2,
			"c": 3,
		}

		result := generic.AnyKey(m, func(k string) bool {
			return strings.HasPrefix(k, "test_");
		})

		t.Log("Should return true if any key satisfies the condition")
		assert.True(t, result)
	})

	t.Run("returns false if no keys satisfy the condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		result := generic.AnyKey(m, func(k string) bool {
			return strings.HasPrefix(k, "test_");
		})

		t.Log("Should return false if no keys satisfy the condition")
		assert.False(t, result)
	})
}

func TestAllValues(t *testing.T) {
	t.Parallel()

	t.Run("returns true if all values satisfy a condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		result := generic.AllValues(m, func(v int) bool {
			return v > 0;
		})

		t.Log("Should return true if all values satisfy the condition")
		assert.True(t, result)
	})

	t.Run("returns false if any value does not satisfy the condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": -1,
			"b": 2,
			"c": 3,
		}

		result := generic.AllValues(m, func(v int) bool {
			return v > 0;
		})

		t.Log("Should return false if any value does not satisfy the condition")
		assert.False(t, result)
	})

	t.Run("handles complex condition", func(t *testing.T) {
		t.Parallel()

		type User struct {
			Name  string
			Admin bool
		}

		m := map[string]User{
			"a": {Name: "Alice", Admin: false},
			"b": {Name: "Bob", Admin: false},
			"c": {Name: "Charlie", Admin: false},
		}
		result := generic.AllValues(m, func(v User) bool {
			return !v.Admin
		})

		t.Log("Should return true if all users are not admins")
		assert.True(t, result)

		m["d"] = User{Name: "Dave", Admin: true}
		result = generic.AllValues(m, func(v User) bool {
			return !v.Admin
		})

		t.Log("Should return false if any user is an admin")
		assert.False(t, result)
	})
}

func TestAnyValue(t *testing.T) {
	t.Parallel()

	t.Run("returns true if any value satisfies a condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": 1,
			"b": -2,
			"c": 3,
		}

		result := generic.AnyValue(m, func(v int) bool {
			return v < 0;
		})

		t.Log("Should return true if any value satisfies the condition")
		assert.True(t, result)
	})

	t.Run("returns false if no values satisfy the condition", func(t *testing.T) {
		t.Parallel()

		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		result := generic.AnyValue(m, func(v int) bool {
			return v < 0;
		})

		t.Log("Should return false if no values satisfy the condition")
		assert.False(t, result)
	})

	t.Run("handles complex condition", func(t *testing.T) {
		t.Parallel()

		type User struct {
			Name  string
			Admin bool
		}

		m := map[string]User{
			"a": {Name: "Alice", Admin: false},
			"b": {Name: "Bob", Admin: true},
			"c": {Name: "Charlie", Admin: false},
		}
		result := generic.AnyValue(m, func(v User) bool {
			return v.Admin
		})

		t.Log("Should return true if any user is an admin")
		assert.True(t, result)

		m["b"] = User{Name: "Bob", Admin: false}
		result = generic.AnyValue(m, func(v User) bool {
			return v.Admin
		})

		t.Log("Should return false if no users are admins")
		assert.False(t, result)
	})
}
