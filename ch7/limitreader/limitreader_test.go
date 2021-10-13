package limitreader

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("Hello! How far will this read?")
	lr := LimitReader(r, 10)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		t.Errorf("there was an error: %v", err)
	}
}
