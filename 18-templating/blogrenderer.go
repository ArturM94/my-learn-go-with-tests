package blogrenderer

import (
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	var tags strings.Builder
	for _, tag := range p.Tags {
		_, err := tags.WriteString(fmt.Sprintf("<li>%s</li>", tag))
		return err
	}
	_, err := fmt.Fprintf(w, `<h1>%s</h1>
<p>%s</p>
Tags: <ul>%s</ul>`, p.Title, p.Description, tags.String())
	return err
}
