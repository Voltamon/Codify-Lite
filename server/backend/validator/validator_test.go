package validator

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestValidator(t *testing.T) {
	g := Goblin(t)

	g.Describe("ValidateName", func() {
		g.It("should pass for valid names", func() {
			err := ValidateName("John Doe")
			g.Assert(err).IsNil()
		})

		g.It("should fail for names with numbers", func() {
			err := ValidateName("John Doe 123")
			g.Assert(err).IsNotNil()
		})

		g.It("should fail for empty names", func() {
			err := ValidateName("")
			g.Assert(err).IsNotNil()
		})
	})
}
