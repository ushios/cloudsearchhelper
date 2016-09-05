package cloudsearchhelper

import (
	"bytes"
	"fmt"
)

// StructuredQueryTokener output query token
// see https://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-compound-queries.html
type StructuredQueryTokener interface {
	QueryString() string
	NotQueryString() string
}

// Near .
type Near struct {
	Distance int
	Boost    int
	Field    string
	Value    string
}

// QueryString make AND or OR query
func (n *Near) QueryString() string {
	var b bytes.Buffer

	// open
	b.WriteString("(near ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", n.Field))

	// ディスタンス
	b.WriteString(fmt.Sprintf(" distance=%d ", n.Distance))

	// Boost
	if n.Boost > 1 {
		b.WriteString(fmt.Sprintf(" boost=%d ", n.Boost))
	}

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", n.Value))

	// close
	b.WriteString(")")

	return b.String()
}

// NotQueryString make not query
func (n *Near) NotQueryString() string {
	var b bytes.Buffer

	// open
	b.WriteString("(near ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", n.Field))

	// ディスタンス
	b.WriteString(fmt.Sprintf(" distance=%d ", n.Distance))

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", n.Value))

	// close
	b.WriteString(")")

	return b.String()
}

// Prefix .
type Prefix struct {
	Boost int
	Field string
	Value string
}

// Phrase .
type Phrase struct {
	Boost int
	Field string
	Value string
}

// Range .
type Range struct {
	Boost int
	Field string
	Value string
}

// Term .
type Term struct {
	Boost int
	Field string
	Value string
}
