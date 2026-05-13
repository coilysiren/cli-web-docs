package layout_test

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/coilysiren/cli-web-docs/layout"
)

// Render a minimal page with title + body. Both cli-web-docs and
// cli-web-ops use this shell so their output is visually consistent.
func ExampleRender() {
	var buf bytes.Buffer
	_ = layout.Render(&buf, layout.Page{
		Title: "Hello",
		Body:  "<p>This is the body.</p>",
	})
	fmt.Println("has title:", strings.Contains(buf.String(), "<title>Hello</title>"))
	fmt.Println("has body :", strings.Contains(buf.String(), "<p>This is the body.</p>"))
	// Output:
	// has title: true
	// has body : true
}

// BuildNavList renders a typed nav-item tree into HTML suitable for the
// Page.Nav field. Children nest under their parent <li>.
func ExampleBuildNavList() {
	html := layout.BuildNavList([]layout.NavItem{
		{Label: "Home", Href: "/index.html"},
		{
			Label: "Group",
			Children: []layout.NavItem{
				{Label: "Item A", Href: "/a.html"},
				{Label: "Item B", Href: "/b.html", Subtitle: "second"},
			},
		},
	})
	fmt.Println(strings.Contains(string(html), `href="/a.html"`))
	// Output: true
}
