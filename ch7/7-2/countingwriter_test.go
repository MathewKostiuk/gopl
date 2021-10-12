package countingwriter

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Hello, world\n")
	fmt.Println(*count)
}
