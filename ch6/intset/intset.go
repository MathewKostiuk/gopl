package intset

import (
	"bytes"
	"fmt"
)

const size = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.

type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/size, uint(x%size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// HasAll reports whether the set contains all of the non-negative values vals.
func (s *IntSet) HasAll(vals ...int) bool {
	has := true

	for _, x := range vals {
		h := s.Has(x)
		if !h {
			has = false
			break
		}
	}
	return has
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/size, uint(x%size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Add all values
func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		s.Add(x)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", size*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		len++
	}
	return len
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	if has := s.Has(x); has {
		word := x / size
		s.words[word] = 0
	}
}

// Remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var y IntSet
	y.words = make([]uint, len(s.words))

	for i, word := range s.words {
		y.words[i] = word
	}
	return &y
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] & (^tword)
		}
	}
}

// SymmetricDifferenceWith returns a set that contains the
// elements present in one set or the other but not both.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	sc := s.Copy()
	// set s to the difference of s and t
	s.DifferenceWith(t)
	// set t to the difference of t and s
	t.DifferenceWith(sc)
	// set s to the union of s and t (differences)
	s.UnionWith(t)
}

// Elems returns a slice containing the elements of the set.
// Suitable for iterating over with a `range` loop
func (s *IntSet) Elems() *[]int {
	var sl []int

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				sl = append(sl, size*i+j)
			}
		}
	}
	return &sl
}
