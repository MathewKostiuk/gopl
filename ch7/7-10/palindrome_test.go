package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(sort.StringSlice([]string{"n", "o", "o", "n"})) {
		t.Errorf("%s should be a palindrome", "\"noon\"")
	}
}
