package service

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestTodo(t *testing.T) {
	g := Goblin(t)
	g.Describe("Todo Request", func() {
		g.It("Should validate description length", func() {
			req := TodoRequest{Description: "Hi"}
			err := req.Validate()
			g.Assert(err).IsNotNil()

			req.Description = "Valid Description"
			err = req.Validate()
			g.Assert(err).IsNil()
		})
	})
}
