package cloudsearchhelper

import (
	"bytes"
	"fmt"
)

// Queryer output query token
// see https://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-compound-queries.html
// syntax https://docs.aws.amazon.com/cloudsearch/latest/developerguide/search-api.html#structured-search-syntax
type Queryer interface {
	QueryString() (string, error)
	NotQueryString() (string, error)
}

// Near .
type Near struct {
	Distance int
	Boost    int
	Field    string
	Value    string
}

// QueryString make AND or OR query
func (n *Near) QueryString() (string, error) {
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

	return b.String(), nil
}

// NotQueryString make not query
func (n *Near) NotQueryString() (string, error) {
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

	return b.String(), nil
}

// Prefix .
type Prefix struct {
	Boost int
	Field string
	Value string
}

// QueryString make AND or OR query
func (p *Prefix) QueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(prefix ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", p.Field))

	// Boost
	if p.Boost > 1 {
		b.WriteString(fmt.Sprintf(" boost=%d ", p.Boost))
	}

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", p.Value))

	// close
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString make not query
func (p *Prefix) NotQueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(prefix ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", p.Field))

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", p.Value))

	// close
	b.WriteString(")")

	return b.String(), nil
}

// Phrase .
type Phrase struct {
	Boost int
	Field string
	Value string
}

// QueryString make AND or OR query
func (p *Phrase) QueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(phrase ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", p.Field))

	// Boost
	if p.Boost > 1 {
		b.WriteString(fmt.Sprintf(" boost=%d ", p.Boost))
	}

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", p.Value))

	// close
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString make not query
func (p *Phrase) NotQueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(phrase ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", p.Field))

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", p.Value))

	// close
	b.WriteString(")")

	return b.String(), nil
}

// Range .
type Range struct {
	Boost int
	Field string
	From  interface{}
	To    interface{}
}

// QueryString make AND or OR query
func (r *Range) QueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(range ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", r.Field))

	// Boost
	if r.Boost > 1 {
		b.WriteString(fmt.Sprintf(" boost=%d ", r.Boost))
	}

	// body
	b.WriteString("[")
	if r.From != nil {
		b.WriteString(fmt.Sprintf("%s", r.From))
	}
	b.WriteString(",")
	if r.To != nil {
		b.WriteString(fmt.Sprintf("%s", r.To))
	}
	b.WriteString("]")

	// close
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString make not query
func (r *Range) NotQueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(range ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", r.Field))

	// body
	b.WriteString(fmt.Sprintf(" [%s, %s] ", r.From, r.To))

	// close
	b.WriteString(")")

	return b.String(), nil
}

// Term .
type Term struct {
	Boost int
	Field string
	Value interface{}
}

// QueryString make AND or OR query
func (t *Term) QueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(term ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", t.Field))

	// Boost
	if t.Boost > 1 {
		b.WriteString(fmt.Sprintf(" boost=%d ", t.Boost))
	}

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", t.Value))

	// close
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString make not query
func (t *Term) NotQueryString() (string, error) {
	var b bytes.Buffer

	// open
	b.WriteString("(term ")

	// 対象フィールド
	b.WriteString(fmt.Sprintf(" field='%s' ", t.Field))

	// body
	b.WriteString(fmt.Sprintf(" '%s' ", t.Value))

	// close
	b.WriteString(")")

	return b.String(), nil
}

func valueString(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	// TODO: convert to string using driver.Valuer

	return "", fmt.Errorf("(%v) is not value", v)
}
