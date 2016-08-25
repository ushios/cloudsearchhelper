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
