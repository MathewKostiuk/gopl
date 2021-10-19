package palindrome

import (
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < s.Len()/2; i, j = i+1, j-1 {
		if !s.Less(i, j) && s.Less(j, i) {
			return false
		}
	}
	return true
}
