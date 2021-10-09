package intset

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))
}

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1000)
	if has := x.Has(1000); !has {
		t.Errorf("the value %d is not in the IntSet", 1000)
	}
}
