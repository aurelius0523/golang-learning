package integers

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("expected %d but got %d", want, got);
	}
}

func ExampleAdd() {
	fmt.Println(Add(2, 2))
	//Output: 4
}
