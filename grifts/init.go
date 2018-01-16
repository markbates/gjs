package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/gjs/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
