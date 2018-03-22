package grifts

import (
	"github.com/TiiQu-Network/claim-verifier-prototype/test/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
