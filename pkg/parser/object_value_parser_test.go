package parser

import (
	"bytes"
	. "github.com/franela/goblin"
	document "github.com/jensneuse/graphql-go-tools/pkg/document"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

func TestObjectValueParser(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parser.parsePeekedObjectValue", func() {

		tests := []struct {
			it        string
			input     string
			expectErr types.GomegaMatcher
			expectVal types.GomegaMatcher
		}{
			{
				it:        "should parse simple object value",
				input:     `{ foo: "bar" }`,
				expectErr: BeNil(),
				expectVal: Equal(document.ObjectValue{
					Val: []document.ObjectField{
						{
							Name: []byte("foo"),
							Value: document.StringValue{
								Val: []byte("bar"),
							},
						},
					},
				}),
			},
			{
				it:        "should parse multiple values",
				input:     `{ foo: "bar" baz: "bat", bas: "bal" anEnum: NUM, smallEnum: numnum }`,
				expectErr: BeNil(),
				expectVal: Equal(document.ObjectValue{
					Val: []document.ObjectField{
						{
							Name: []byte("foo"),
							Value: document.StringValue{
								Val: []byte("bar"),
							},
						},
						{
							Name: []byte("baz"),
							Value: document.StringValue{
								Val: []byte("bat"),
							},
						},
						{
							Name: []byte("bas"),
							Value: document.StringValue{
								Val: []byte("bal"),
							},
						},
						{
							Name: []byte("anEnum"),
							Value: document.EnumValue{
								Name: []byte("NUM"),
							},
						},
						{
							Name: []byte("smallEnum"),
							Value: document.EnumValue{
								Name: []byte("numnum"),
							},
						},
					},
				}),
			},
			{
				it:        "should parse nested object value",
				input:     `{ foo: { bar: "baz" } }`,
				expectErr: BeNil(),
				expectVal: Equal(document.ObjectValue{
					Val: []document.ObjectField{
						{
							Name: []byte("foo"),
							Value: document.ObjectValue{
								Val: []document.ObjectField{
									{
										Name: []byte("bar"),
										Value: document.StringValue{
											Val: []byte("baz"),
										},
									},
								},
							},
						},
					},
				}),
			},
			{
				it: "should parse nested object value across multiple lines",
				input: `{foo	:
	{
		bar: "baz"
	}
}`,
				expectErr: BeNil(),
				expectVal: Equal(document.ObjectValue{
					Val: []document.ObjectField{
						{
							Name: []byte("foo"),
							Value: document.ObjectValue{
								Val: []document.ObjectField{
									{
										Name: []byte("bar"),
										Value: document.StringValue{
											Val: []byte("baz"),
										},
									},
								},
							},
						},
					},
				}),
			},
		}

		for _, test := range tests {
			test := test

			g.It(test.it, func() {

				reader := bytes.NewReader([]byte(test.input))
				parser := NewParser()
				parser.l.SetInput(reader)

				val, err := parser.parsePeekedObjectValue()
				Expect(err).To(test.expectErr)
				Expect(val).To(test.expectVal)
			})
		}
	})
}
