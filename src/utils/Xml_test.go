package utils

import (
	"testing"
	"fmt"
)

func TestTagParams(t *testing.T) {
	const html = `
<search id="mysearch">
</search>
`
	fmt.Println(Xml.TagAtrrs(html, "search", 0))
}
