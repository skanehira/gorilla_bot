package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
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

// ReadFile read data from file
func ReadFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		msg := fmt.Sprintf("read message file is failed %s", err)
		log.Printf("[%s] %s", TimeNow(), msg)
		return make([]byte, 0), errors.New(msg)
	}
	return data, nil
}

// StructToMap struct convert to map
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}
