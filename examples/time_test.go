package examples

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ushios/cloudsearchhelper"
)

type Event struct {
	Name    string
	StartAt cloudsearchhelper.Time
}

func ExampleEvent() {
	now, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")

	e := Event{
		Name:    "This is some event",
		StartAt: cloudsearchhelper.Time(now),
	}

	j, _ := json.Marshal(e)

	fmt.Println(string(j))
	// Output: {"Name":"This is some event","StartAt":"2006-01-02T08:04:05Z"}
}
