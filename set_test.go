package generic_test

// This file generated with Claude 3.7 Sonnet

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/singlestore-labs/generic"
)

func TestToSet(t *testing.T) {
	t.Parallel()

	t.Run("converts slice to set", func(t *testing.T) {
		t.Parallel()

		slice := []int{1, 2, 3, 4, 5}
		set := generic.ToSet(slice)

		t.Log("Should create set with all elements from slice")
		assert.Len(t, set, 5)

		for _, item := range slice {
			_, exists := set[item]
			assert.True(t, exists, "Set should contain %v", item)
		}
	})

	t.Run("handles duplicates", func(t *testing.T) {
		t.Parallel()

		slice := []string{"a", "b", "a", "c", "b"}
		set := generic.ToSet(slice)

		t.Log("Should create set with unique elements from slice")
		assert.Len(t, set, 3)

		expected := map[string]struct{}{
			"a": {},
			"b": {},
			"c": {},
		}

		assert.Equal(t, expected, set)
	})

	t.Run("handles empty slice", func(t *testing.T) {
		t.Parallel()

		slice := []int{}
		set := generic.ToSet(slice)

		t.Log("Should create empty set from empty slice")
		assert.Empty(t, set)
	})
}
