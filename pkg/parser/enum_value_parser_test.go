package parser

import (
	"bytes"
	. "github.com/franela/goblin"
	"github.com/jensneuse/graphql-go-tools/pkg/document"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

func TestEnumValueParser(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parser.parsePeekedEnumValue", func() {

		tests := []struct {
			it         string
			input      string
			expectErr  types.GomegaMatcher
			expectName types.GomegaMatcher
		}{
			{
				it:         "should parse MY_ENUM",
				input:      "MY_ENUM",
				expectErr:  BeNil(),
				expectName: Equal(document.ByteSlice("MY_ENUM")),
			},
		}

		for _, test := range tests {
			test := test

			g.It(test.it, func() {

				reader := bytes.NewReader([]byte(test.input))
				parser := NewParser()
				parser.l.SetInput(reader)

				val, err := parser.parsePeekedEnumValue()
				Expect(err).To(test.expectErr)
				Expect(val.Name).To(test.expectName)
			})
		}
	})
}
