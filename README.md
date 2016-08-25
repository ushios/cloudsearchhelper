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


### Time

Cloudsearch date field's format using RFC3339 ([see detail](https://docs.aws.amazon.com/ja_jp/cloudsearch/latest/developerguide/configuring-index-fields.html)).
This `Time` return UTC RFC3339 string when MarshalJSON.

```go
type Event struct {
	Name  string
	Start cloudsearchhelper.Time
}

e := Event{
 	Name : "game show",
 	Start : time.Now(),
}

j, _ := json.MarshalJSON(e)
fmt.Println(string(j)) // The Start json value(UTC) -> "2006-01-02T15:04:05Z"
```
