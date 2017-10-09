package validate

import (
	"testing"
)

func TestSanitizeRules(t *testing.T) {
	set := []struct {
		input  string
		output []string
	}{
		{
			input:  "apples, oranges , bananas",
			output: []string{"apples", "oranges", "bananas"},
		},
		{
			input:  "apples, oranges , bananas, peaches='with cream'",
			output: []string{"apples", "oranges", "bananas", "peaches='with cream'"},
		},
	}

	for i := range set {
		filt := SanitizeFilters(set[i].input)
		if !StringsSliceEqual(filt, set[i].output) {
			t.Errorf("expected output like %v, but got %v", set[i].output, filt)
		}
	}
}

func TestSliceHelpers(t *testing.T) {
	set := []struct {
		A, B  []string
		Equal bool
	}{
		{
			A:     []string{"apples", "bananas", "oranges", "mango"},
			B:     []string{"oranges", "bananas", "apples", "mango"},
			Equal: true,
		},
		{
			A:     []string{"apples", "bananas", "oranges", "mango"},
			B:     []string{"waffles"},
			Equal: false,
		},
	}
	for i := range set {
		got := StringsSliceEqual(set[i].A, set[i].B)
		if got != set[i].Equal {
			t.Errorf("expected %d A and B comparison to be %t, but got %t", i, set[i].Equal, got)
		}
	}
}
