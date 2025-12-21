package slices_test

import (
	"strconv"
	"testing"

	slicespkg "github.com/joseMarciano/crypto-manager/pkg/slices"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Run("should map strings to their lengths", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		expected := []int{5, 6, 6}

		result := slicespkg.Map(input, func(s string) int { return len(s) })

		require.Len(t, expected, len(result))
		require.Equal(t, expected, result)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		input := []int{}
		expected := []string{}

		result := slicespkg.Map(input, func(i int) string { return strconv.Itoa(i) })

		require.Len(t, expected, len(result))
		require.Equal(t, expected, result)
	})

	t.Run("should map custom structs to field values", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		input := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		expected := []string{"Alice", "Bob", "Charlie"}

		result := slicespkg.Map(input, func(p Person) string { return p.Name })

		require.Len(t, expected, len(result))
		require.Equal(t, expected, result)
	})

	t.Run("should map pointers to values", func(t *testing.T) {
		val1, val2, val3 := 10, 20, 30
		input := []*int{&val1, &val2, &val3}
		expected := []int{10, 20, 30}

		result := slicespkg.Map(input, func(p *int) int { return *p })

		require.Len(t, expected, len(result))
		require.Equal(t, expected, result)
	})

	t.Run("should map interfaces to concrete types", func(t *testing.T) {
		input := []interface{}{1, 2, 3, 4}
		expected := []int{1, 2, 3, 4}

		result := slicespkg.Map(input, func(i interface{}) int { return i.(int) })

		require.Len(t, expected, len(result))
		require.Equal(t, expected, result)
	})
}
