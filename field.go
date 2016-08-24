package cloudsearchhelper

// See more - https://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html

type (
	// FieldType is type of cloudsearch fields
	FieldType int

	// FieldOptions are field option list
	FieldOptions struct {
		HighlightEnabled bool
		FacetEnabled     bool
		ReturnEnabled    bool
		SearchEnabled    bool
		SortEnabled      bool
	}
)

const (
	Date FieldType = iota
	DateArray
	Double
	DoubleArray
	Int
	IntArray
	LatLon
	Literal
	LiteralArray
	Text
	TextArray
)
