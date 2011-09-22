package ssgen

import (
	"testing"
	"bytes"
)

func Test_GetNewFilename_WithExt(t *testing.T) {
	m := NewMarkdownToHtml()
	for _, ext := range m.supported_exts {
		assertEqual(t, m.GetNewFilename("test"+ext), "test.html")
	}
}

func Test_GetNewFilename_WithoutExt(t *testing.T) {
	m := NewMarkdownToHtml()
	assertEqual(t, m.GetNewFilename("test"), "test.html")
}

func Test_Supports(t *testing.T) {
	m := NewMarkdownToHtml()
	assertEqual(t, m.Supports("test.md"), true)
	assertEqual(t, m.Supports("test.mkdwn"), true)
}
func Test_SimpleConversion(t *testing.T) {
	from := bytes.NewBufferString("hello")
	to := bytes.NewBufferString("")
	var m Converter

	m = NewMarkdownToHtml()
	m.Convert(from, to)

	assertEqual(t, to.String(), "<p>hello</p>\n")
}
