package main

import (
	"strings"

	"github.com/gobuffalo/plush"
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
)

func main() {
	ctx := plush.NewContext()
	ctx.Set("name", "Ringo")
	ctx.Set("mCall", func(help plush.HelperContext) (string, error) {
		n, ok := help.Value("name").(string)
		if !ok {
			return "", errors.Errorf("expected name to be a string, got %T", help.Value("name"))
		}
		return strings.ToUpper(n), nil
	})
	s, err := plush.Render(t, ctx)
	if err != nil {
		panic(err)
	}
	js.Global.Get("document").Call("write", s)
}

const t = `
Hello <%= name %>

<%= mCall() %>
`
