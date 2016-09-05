ushios/cloudsearchhelper
=========================

[![Build Status](https://travis-ci.org/ushios/cloudsearchhelper.svg?branch=master)](https://travis-ci.org/ushios/cloudsearchhelper)
[![Coverage Status](https://coveralls.io/repos/ushios/cloudsearchhelper/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/cloudsearchhelper?branch=master)

cloudsearch helper


Installation
=============

```bash
$ go get github.com/ushios/cloudsearchhelper
```


Usage
=====

- [Time](#time)
- [StructuredQuery](#StructuredQuery)


### Time

when create documents

Cloudsearch date field's format using RFC3339 ([see detail](https://docs.aws.amazon.com/ja_jp/cloudsearch/latest/developerguide/configuring-index-fields.html)).
This `Time` return UTC RFC3339 string when MarshalJSON.

example [here](https://github.com/ushios/cloudsearchhelper/blob/master/examples/time_test.go#L16-L28)

```go
now, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")

e := Event{
	Name:    "This is some event",
	StartAt: cloudsearchhelper.Time(now),
}

j, _ := json.Marshal(e)

fmt.Println(string(j))
// Output: {"Name":"This is some event","StartAt":"2006-01-02T08:04:05Z"}
```

### StructuredQuery


Sorry time.Time not supported now x(

example [here](https://github.com/ushios/cloudsearchhelper/blob/master/examples/query_test.go#L9-L42)

```go
category := &cloudsearchhelper.Prefix{
	Field: "category",
	Value: "movie",
}

title := &cloudsearchhelper.Phrase{
	Field: "title",
	Value: "star",
}

and := cloudsearchhelper.And(&[]cloudsearchhelper.Queryer{
	category,
	title,
	or,
})

fmt.Println(and.QueryString())
// Output: (and (prefix  field='category'  'movie' )(phrase  field='title'  'star' )))

```
