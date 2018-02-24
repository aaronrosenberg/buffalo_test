package grifts

import (
	"github.com/aaronrosenberg/buffalo_test/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
