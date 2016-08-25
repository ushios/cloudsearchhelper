package cloudsearchhelper

import (
	"bytes"
	"time"
)

const (
	// DateFormat for cloudsearch date type
	DateFormat = time.RFC3339
)

// Time for cloudsearch indices
type Time time.Time

// MarshalJSON for json.Marshaler
func (t Time) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	tm := time.Time(t)
	buf.WriteString("\"")
	buf.WriteString(tm.UTC().Format(DateFormat))
	buf.WriteString("\"")
	return buf.Bytes(), nil
}
