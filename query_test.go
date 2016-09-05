package cloudsearchhelper

import "testing"

func TestQuery(t *testing.T) {
	test := func(query StructuredQueryTokener, e string) {
		if e != query.QueryString() {
			t.Errorf("%v query expected (%s) but (%s)", query, e, query.QueryString())
		}
	}

	n := Near{
		Distance: 3,
		Boost:    1,
		Field:    "title",
		Value:    "Teenage Ninja Mutant Turtles",
	}

	p := Prefix{
		Boost: 1,
		Field: "title",
		Value: "Tennage Ninja Mutant Turtles",
	}

	n2 := Near{
		Distance: 1,
		Boost:    1,
		Field:    "title",
		Value:    "Star Wars",
	}

	a := And(&[]StructuredQueryTokener{&n2})

	o := Or(&[]StructuredQueryTokener{&n, &p, a})

	test(o, `(or (near  field='title'  distance=3  'Teenage Ninja Mutant Turtles' )(prefix  field='title'  'Tennage Ninja Mutant Turtles' )(and (near  field='title'  distance=1  'Star Wars' )))`)

}

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
