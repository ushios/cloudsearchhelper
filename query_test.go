package cloudsearchhelper

import "testing"

func TestQuery(t *testing.T) {
	test := func(query Queryer, e string) {
		qs, err := query.QueryString()
		if err != nil {
			t.Fatal(err)
		}
		if e != qs {
			t.Errorf("%v query expected (%s) but (%s)", query, e, qs)
		}
	}

	title1 := Prefix{
		Boost: 1,
		Field: "title",
		Value: "new",
	}

	title2 := Prefix{
		Boost: 1,
		Field: "title",
		Value: "word",
	}

	body1 := Prefix{
		Boost: 1,
		Field: "body",
		Value: "question",
	}

	o := Or(&[]Queryer{&title1, &title2})
	a := And(&[]Queryer{&body1, o})

	test(a, `(and (prefix  field='body'  'question' )(or (prefix  field='title'  'new' )(prefix  field='title'  'word' )))`)

}

func TestAndQuery(t *testing.T) {
	test := func(tokens []Queryer, e string) {
		a := And(&tokens)

		qs, err := a.QueryString()
		if err != nil {
			t.Fatal(err)
		}
		if e != qs {
			t.Errorf("%v's query expected (%s) but (%s)", tokens, e, qs)
		}
	}

	test([]Queryer{
		&Near{
			Distance: 1,
			Boost:    1,
			Field:    "hoge",
			Value:    "fuga",
		},
	}, `(and (near  field='hoge'  distance=1  'fuga' ))`)

	test([]Queryer{
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

func TestNotQuery(t *testing.T) {
	test := func(tokens []Queryer, e string) {
		n := Not(&tokens)
		qs, err := n.QueryString()
		if err != nil {
			t.Fatal(err)
		}
		if e != qs {
			t.Errorf("%v's query expected (%s) but (%s)", tokens, e, qs)
		}
	}

	test([]Queryer{
		&Phrase{
			Field: "title",
			Value: "Star Wars",
		},
	}, `(not (phrase  field='title'  'Star Wars' ))`)
}
