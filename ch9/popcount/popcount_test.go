package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	fmt.Println(PopCount(uint64(15)))
	fmt.Println(PopCount(uint64(109)))
	fmt.Println(PopCount(uint64(20)))

}
