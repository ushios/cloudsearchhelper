package cloudsearchhelper

import "bytes"

// AndQuery is token for and clause
type AndQuery struct {
	tokens *[]Queryer
}

// QueryString .
func (a *AndQuery) QueryString() (string, error) {
	var b bytes.Buffer

	if a.tokens == nil || len(*a.tokens) == 0 {
		return "", nil
	}

	b.WriteString("(and ")
	for _, t := range *a.tokens {
		qs, err := t.QueryString()
		if err != nil {
			return "", err
		}
		b.WriteString(qs)
	}
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString for not
func (a *AndQuery) NotQueryString() (string, error) {
	return a.QueryString()
}

// OrQuery .
type OrQuery struct {
	tokens *[]Queryer
}

// QueryString .
func (o *OrQuery) QueryString() (string, error) {
	var b bytes.Buffer

	if o.tokens == nil || len(*o.tokens) == 0 {
		return "", nil
	}

	b.WriteString("(or ")
	for _, t := range *o.tokens {
		qs, err := t.QueryString()
		if err != nil {
			return "", err
		}
		b.WriteString(qs)
	}
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString for not
func (o *OrQuery) NotQueryString() (string, error) {
	return o.QueryString()
}

// NotQuery .
type NotQuery struct {
	tokens *[]Queryer
}

// QueryString .
func (n *NotQuery) QueryString() (string, error) {
	var b bytes.Buffer

	if n.tokens == nil || len(*n.tokens) == 0 {
		return "", nil
	}

	b.WriteString("(not ")
	for _, t := range *n.tokens {
		qs, err := t.NotQueryString()
		if err != nil {
			return "", err
		}
		b.WriteString(qs)
	}
	b.WriteString(")")

	return b.String(), nil
}

// NotQueryString for not
func (n *NotQuery) NotQueryString() (string, error) {
	return n.QueryString()
}

// And .
func And(tokens *[]Queryer) *AndQuery {
	return &AndQuery{
		tokens: tokens,
	}
}

// Or .
func Or(tokens *[]Queryer) *OrQuery {
	return &OrQuery{
		tokens: tokens,
	}
}

// Not .
func Not(tokens *[]Queryer) *NotQuery {
	return &NotQuery{
		tokens: tokens,
	}
}
