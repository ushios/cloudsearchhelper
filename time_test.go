package cloudsearchhelper

import (
	"testing"
	"time"
)

func TestTimeMarshalJSON(t *testing.T) {
	test := func(timeStr string, j string) {
		tm, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			t.Fatal(err)
		}

		str, err := tm.MarshalJSON()
		if err != nil {
			t.Error(err)
		}

		if string(str) != j {
			t.Errorf("(%s) marshaljson expected(%s) but (%s)", timeStr, j, str)
		}
	}

	test("2016-10-09T10:00:00Z09:00", `{"2016-10-09T01:00:00Z"}`)
}
