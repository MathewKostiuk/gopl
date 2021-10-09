package intset

import (
	"testing"
)

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1000)
	if has := x.Has(1000); !has {
		t.Errorf("the value %d is not in the IntSet", 1000)
	}
}

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	has1 := x.Has(1)
	if !has1 {
		t.Errorf("the intset does not contain %d.", 1)
	}
	has2 := x.Has(2)
	if !has2 {
		t.Errorf("the intset does not contain %d.", 2)
	}
	has3 := x.Has(3)
	if !has3 {
		t.Errorf("the intset does not contain %d.", 3)
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)

	h := x.HasAll(1, 2, 3)
	if !h {
		t.Error("the intset does not contain all values")
	}
}

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(10)
	x.Add(100)
	x.Add(1000)

	len := x.Len()
	if len != 3 {
		t.Errorf("length of IntSet does not =%d", 3)
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(10)
	x.Add(1001)

	x.Remove(10)

	if has := x.Has(10); has {
		t.Errorf("the value %d was not removed from the IntSet", 10)
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(10)
	x.Add(100)
	x.Add(1000)

	x.Clear()
	len := x.Len()
	if len != 0 {
		t.Error("the IntSet was not cleared")
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(10)
	x.Add(100)
	x.Add(1000)

	y := x.Copy()

	has1 := y.Has(10)
	has2 := y.Has(100)
	has3 := y.Has(1000)
	if !has1 && !has2 && !has3 {
		t.Errorf("the copied intset does not have %d, %d, %d.", 10, 100, 1000)
	}
}

func TestUnion(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)

	hasNine, hasOneTwoThree := x.Has(9), x.Has(123)
	if !hasNine && !hasOneTwoThree {
		t.Errorf("the intset does not contain %d and %d", 9, 123)
	}
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(2, 3, 5)
	x.IntersectWith(&y)

	has := x.HasAll(2, 3)
	if !has {
		t.Error("x does not have all intersecting values")
	}

	has1 := x.Has(1)
	if has1 {
		t.Errorf("x still has this element: %d", 1)
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(2, 3, 4)
	x.DifferenceWith(&y)

	has := x.Has(2)
	if has {
		t.Errorf("x still has %d", 2)
	}
	has2 := x.Has(3)
	if has2 {
		t.Errorf("x still has %d", 3)
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(4, 5, 6)
	x.SymmetricDifferenceWith(&y)

	has := x.HasAll(1, 2, 3, 4, 5, 6)
	if !has {
		t.Error("z does not have all of the required elements")
	}
}

func TestElems(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)

	s := x.Elems()

	for i, el := range *s {
		if i == 0 && el != 1 {
			t.Errorf("element is missing: %d", el)
		}
		if i == 1 && el != 2 {
			t.Errorf("element is missing: %d", el)
		}
		if i == 2 && el != 3 {
			t.Errorf("element is missing: %d", el)
		}
	}
}
