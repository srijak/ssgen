## ssgen: Static site generator

Given a directory, convert to html and write out to the target directory.
Anything that can't be converted (no applicable Converter) is simply copied over.


#### NOTE: the target directory is replaced with new output so shouldn't contain anything you mind losing.

### Installation:
```
goinstall github.com/srijak/ssgen
```

### Usage:
```go
import "github.com/srijak/ssgen"
.
.

// initialize the markdown to html converter
m := ssgen.NewMarkdownToHtml()
src_dir := "./a"
target_dir := "./a_web_root"

// give it the:
//   src_dir    : directory with files in original format.
//   target_dir: directory to publish to. *NOTE* this is deleted and rewritten so make sure you can loose it.
//   []Converter: array of converters (you can write your own as long as they implement the Converter interface)
ssgen.Convert(src_dir, target_dir, []Converter{m})

```

