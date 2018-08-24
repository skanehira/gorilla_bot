package common

import (
	"bytes"
	"encoding/json"
	"time"
)

// FormatStringJoin format indent string join
func FormatStringJoin(j string) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(j), "", "  ")
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// TimeNow Get formated time now
func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
