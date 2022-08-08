package general

import (
	"reflect" // package to compare any two variables
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 6})
	expected := []int{3, 6}

	if !reflect.DeepEqual(got, expected) { // can't use equality with slices, use DeepEqual() method to compare
		t.Errorf("got %d expected %d", got, expected)
	}
}
