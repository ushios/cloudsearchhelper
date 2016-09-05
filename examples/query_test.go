package examples

import (
	"fmt"

	"github.com/ushios/cloudsearchhelper"
)

func ExampleQuery() {
	category := &cloudsearchhelper.Prefix{
		Field: "category",
		Value: "movie",
	}

	title := &cloudsearchhelper.Phrase{
		Field: "title",
		Value: "star",
	}

	genre1 := &cloudsearchhelper.Term{
		Field: "genre",
		Value: "SF",
	}

	genre2 := &cloudsearchhelper.Term{
		Field: "genre",
		Value: "battle",
	}

	or := cloudsearchhelper.Or(&[]cloudsearchhelper.Queryer{
		genre1,
		genre2,
	})

	and := cloudsearchhelper.And(&[]cloudsearchhelper.Queryer{
		category,
		title,
		or,
	})
	fmt.Println(and.QueryString())
	// Output: (and (prefix  field='category'  'movie' )(phrase  field='title'  'star' )(or (term  field='genre'  'SF' )(term  field='genre'  'battle' )))
}
