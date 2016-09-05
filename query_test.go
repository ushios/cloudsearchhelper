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

	test([]StructuredQueryTokener{
		&Near{
			Distance: 1,
			Boost:    1,
			Field:    "price",
			Value:    "1000",
		},
		&Near{
			Distance: 2,
			Boost:    5,
			Field:    "title",
			Value:    "シン・ゴジラ",
		},
	}, `(and (near  field='price'  distance=1  '1000' )(near  field='title'  distance=2  boost=5  'シン・ゴジラ' ))`)
}
