package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterValuesInKeys(t *testing.T) {
	integerMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}

	type testCase struct {
		name         string
		keysToFilter []int
		want         []int
	}

	tests := []testCase{
		{
			name:         "When filter value for all keys, Then should return all values",
			keysToFilter: []int{1, 2, 3, 4},
			want:         []int{1, 2, 3, 3},
		},
		{
			name:         "When filter with empty keys to filter, Then always return empty array",
			keysToFilter: []int{},
			want:         []int{},
		},
		{
			name:         "When filter some keys, Then return values for the keys given to filter",
			keysToFilter: []int{2, 4},
			want:         []int{2, 4},
		},
		{
			name:         "When filter with keys that aren't present in map, Then ignore non present",
			keysToFilter: []int{10},
			want:         []int{},
		},
		{
			name:         "Filter keys contains both present and non present, then return only present keys",
			keysToFilter: []int{3, 10},
			want:         []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterValuesInKeys(integerMap, tt.keysToFilter...)

			assert.Len(t, got, len(tt.want))

			for _, want := range tt.want {
				assert.Contains(t, got, want)
			}
		})
	}
}
