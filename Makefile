# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=github.com/srijak/configo
GOFILES=\
	converter.go\
	markdown_to_html.go\
	testutils.go\
	ssgen.go\

include $(GOROOT)/src/Make.pkg
