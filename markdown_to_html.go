package ssgen

import (
	"os"
	"io"
	"io/ioutil"
	"goskirt"
	"path"
	"strings"
)

type MarkdownToHtml struct {
	gs             goskirt.Goskirt
	supported_exts []string
	converted_ext  string
}

func NewMarkdownToHtml() *MarkdownToHtml {
	m := new(MarkdownToHtml)
	m.gs = goskirt.Goskirt{
		goskirt.EXT_AUTOLINK | goskirt.EXT_STRIKETHROUGH,
		goskirt.HTML_SMARTYPANTS | goskirt.HTML_USE_XHTML,
	}
	m.supported_exts = []string{".md", ".mkdwn"}
	m.converted_ext = ".html"
	return m
}

func (m MarkdownToHtml) Convert(r io.Reader, w io.Writer) (int, os.Error) {
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return -1, err
	}
	return m.gs.WriteHTML(w, contents)
}

func (m MarkdownToHtml) Supports(filename string) bool {
	ext := path.Ext(filename)
	for _, e := range m.supported_exts {
		if ext == e {
			return true
		}
	}
	return false
}

func (m MarkdownToHtml) GetNewFilename(filename string) string {
	ext := path.Ext(filename)
	if ext != "" {
		idx := strings.LastIndex(filename, ext)
		return filename[0:idx] + m.converted_ext
	}
	return filename + m.converted_ext
}
