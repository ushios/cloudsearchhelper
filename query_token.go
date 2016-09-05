package cloudsearchhelper

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

// Queryer output query token
// see https://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching-compound-queries.html
// syntax https://docs.aws.amazon.com/cloudsearch/latest/developerguide/search-api.html#structured-search-syntax
type Queryer interface {
	QueryString() (string, error)
	NotQueryString() (string, error)
}

type timeFormatter interface {
	Format(string) string
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
		s, err := valueString(r.From)
		if err != nil {
			return "", err
		}
		b.WriteString(fmt.Sprintf("%s", s))
	}
	b.WriteString(",")
	if r.To != nil {
		s, err := valueString(r.To)
		if err != nil {
			return "", err
		}
		b.WriteString(fmt.Sprintf("%s", s))
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
	s, err := valueString(t.Value)
	if err != nil {
		return "", err
	}
	b.WriteString(fmt.Sprintf(" %s", s))

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

	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%s", fmt.Sprint(v)), nil
	case float32, float64:
		f := v.(float64)
		return fmt.Sprintf("%f", f), nil
	case string:
		s := v.(string)
		return fmt.Sprintf("'%s'", s), nil
	case time.Time:
		t := v.(timeFormatter)
		return fmt.Sprintf("'%s'", t.Format(DateFormat)), nil

	}

	return "", fmt.Errorf("(%s) is not supported", reflect.TypeOf(v))
}
