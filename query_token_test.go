package cloudsearchhelper

import "testing"

func TestNearQueryString(t *testing.T) {
	test := func(n Near, e string) {
		q := n.QueryString()

		if q != e {
			t.Errorf("%v query string expected (%s) but (%s)", n, e, q)
		}
	}

	test(Near{
		Distance: 4,
		Boost:    1,
		Field:    "title",
		Value:    "some title",
	}, `(near  field='title'  distance=4  'some title' )`)

	test(Near{
		Distance: 4,
		Boost:    2,
		Field:    "title",
		Value:    "some title",
	}, `(near  field='title'  distance=4  boost=2  'some title' )`)
}

func TestNearNotQueryString(t *testing.T) {
	test := func(n Near, e string) {
		q := n.NotQueryString()

		if q != e {
			t.Errorf("%v query string expected (%s) but (%s)", n, e, q)
		}
	}

	test(Near{
		Distance: 4,
		Boost:    2,
		Field:    "title",
		Value:    "some title",
	}, `(near  field='title'  distance=4  'some title' )`)
}

func TestPrefixQueryString(t *testing.T) {
	test := func(p Prefix, e string) {
		q := p.QueryString()

		if q != e {
			t.Errorf("%v query string expected (%s) but (%s)", p, e, q)
		}
	}

	test(Prefix{
		Boost: 1,
		Field: "title",
		Value: "some title",
	}, `(prefix  field='title'  'some title' )`)

	test(Prefix{
		Boost: 2,
		Field: "title",
		Value: "some title",
	}, `(prefix  field='title'  boost=2  'some title' )`)
}

func TestPrefixNotQueryString(t *testing.T) {
	test := func(p Prefix, e string) {
		q := p.NotQueryString()

		if q != e {
			t.Errorf("%v query string expected (%s) but (%s)", p, e, q)
		}
	}

	test(Prefix{
		Boost: 2,
		Field: "title",
		Value: "some title",
	}, `(prefix  field='title'  'some title' )`)
}

func TestPhraseQueryString(t *testing.T) {
	test := func(p Phrase, e string) {
		q := p.QueryString()

		if q != e {
			t.Errorf("%v query string expected (%s) but (%s)", p, e, q)
		}
	}

	test(Phrase{
		Boost: 1,
		Field: "title",
		Value: "some title",
	}, `(phrase  field='title'  'some title' )`)

	test(Phrase{
		Boost: 2,
		Field: "title",
		Value: "some title",
	}, `(phrase  field='title'  boost=2  'some title' )`)
}

func TestPhraseNotQueryString(t *testing.T) {
	test := func(p Phrase, e string) {
		q := p.NotQueryString()

		if q != e {
			t.Errorf("%v query string expected (%s) but (%s)", p, e, q)
		}
	}

	test(Phrase{
		Boost: 2,
		Field: "title",
		Value: "some title",
	}, `(phrase  field='title'  'some title' )`)
}
