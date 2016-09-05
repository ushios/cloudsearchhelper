package cloudsearchhelper

import "bytes"

// AndQuery is token for and clause
type AndQuery struct {
	tokens *[]StructuredQueryTokener
}

// QueryString .
func (a *AndQuery) QueryString() string {
	var b bytes.Buffer

	if a.tokens == nil || len(*a.tokens) == 0 {
		return ""
	}

	b.WriteString("(and ")
	for _, t := range *a.tokens {
		b.WriteString(t.QueryString())
	}
	b.WriteString(")")

	return b.String()
}

// NotQueryString for not
func (a *AndQuery) NotQueryString() string {
	return a.QueryString()
}

// OrQuery .
type OrQuery struct {
	tokens *[]StructuredQueryTokener
}

// QueryString .
func (o *OrQuery) QueryString() string {
	var b bytes.Buffer

	if o.tokens == nil || len(*o.tokens) == 0 {
		return ""
	}

	b.WriteString("(or ")
	for _, t := range *o.tokens {
		b.WriteString(t.QueryString())
	}
	b.WriteString(")")

	return b.String()
}

// NotQueryString for not
func (o *OrQuery) NotQueryString() string {
	return o.QueryString()
}

// NotQuery .
type NotQuery struct {
	tokens *[]StructuredQueryTokener
}

// QueryString .
func (n *NotQuery) QueryString() string {
	var b bytes.Buffer

	if n.tokens == nil || len(*n.tokens) == 0 {
		return ""
	}

	b.WriteString("(not ")
	for _, t := range *n.tokens {
		b.WriteString(t.NotQueryString())
	}
	b.WriteString(")")

	return b.String()
}

// NotQueryString for not
func (n *NotQuery) NotQueryString() string {
	return n.QueryString()
}

// And .
func And(tokens *[]StructuredQueryTokener) *AndQuery {
	return &AndQuery{
		tokens: tokens,
	}
}

// Or .
func Or(tokens *[]StructuredQueryTokener) *OrQuery {
	return &OrQuery{
		tokens: tokens,
	}
}

// Not .
func Not(tokens *[]StructuredQueryTokener) *NotQuery {
	return &NotQuery{
		tokens: tokens,
	}
}
