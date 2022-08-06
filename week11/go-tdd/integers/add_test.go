package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 4)
	expected := 6

	if sum != expected {
		t.Errorf("expected %d got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 0)
	fmt.Println(sum)
	// Output: 1
}
