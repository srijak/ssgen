package ssgen

import (
	"io"
	"os"
)

type Converter interface {
	// converts contents of r and writes to w
	Convert(r io.Reader, w io.Writer) (int, os.Error)

	// returns true is this converter supports the given file.
	Supports(filename string) bool

	// returns the filename that shoudl be 
	// given to the file after the convesion.
	GetNewFilename(filename string) string
}
