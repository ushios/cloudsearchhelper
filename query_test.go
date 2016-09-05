package cloudsearchhelper

import "testing"

func TestAndQuery(t *testing.T) {
	test := func(tokens []StructuredQueryTokener, e string) {
		a := And(&tokens)

		if e != a.QueryString() {
			t.Errorf("%v's query expected (%s) but (%s)", tokens, e, a.QueryString())
		}
	}

	test([]StructuredQueryTokener{
		&Near{
			Distance: 1,
			Boost:    1,
			Field:    "hoge",
			Value:    "fuga",
		},
	}, `(and (near  field='hoge'  distance=1  'fuga' ))`)
}
