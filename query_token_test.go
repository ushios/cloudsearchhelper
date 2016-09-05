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
