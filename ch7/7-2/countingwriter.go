package countingwriter

import "io"

type ByteCounter struct {
	writer io.Writer
	count  int64
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	n, err := b.writer.Write(p)
	b.count += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{w, 0}
	return &bc, &bc.count
}
