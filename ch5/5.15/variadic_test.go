package variadic

import "testing"

var tests = []struct {
	values []int
	result map[string]int
}{
	{[]int{1, 4, 25, 999}, map[string]int{
		"max": 999,
		"min": 1,
	}},
	{[]int{2, 3, 1, 0}, map[string]int{
		"max": 3,
		"min": 0,
	}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, map[string]int{
		"max": 10,
		"min": 1,
	}},
	{[]int{}, map[string]int{
		"max": 0,
		"min": 0,
	}},
}

func TestMax(t *testing.T) {
	for _, test := range tests {
		max := Max(test.values...)
		if max != test.result["max"] {
			t.Errorf(`Max(%d) = %d is false, result=%d`, test.values, test.result["max"], max)
		}
	}
}

func TestMin(t *testing.T) {
	for _, test := range tests {
		min := Min(test.values...)
		if min != test.result["min"] {
			t.Errorf(`Min(%d) = %d is false, result=%d`, test.values, test.result["min"], min)
		}
	}
}
