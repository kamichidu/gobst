package funcs

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
)

func jsonStringify(v interface{}) string {
	var buffer bytes.Buffer
	je := json.NewEncoder(&buffer)
	if err := je.Encode(v); err != nil {
		panic(pkgName + ": unable to encode to json: " + err.Error())
	}
	return buffer.String()
}

func jsonPretty(v interface{}) string {
	var buffer bytes.Buffer
	je := json.NewEncoder(&buffer)
	je.SetIndent("", "  ")
	if err := je.Encode(v); err != nil {
		panic(pkgName + ": unable to encode to json: " + err.Error())
	}
	return buffer.String()
}

func jsonParse(v interface{}) interface{} {
	var r io.Reader
	switch v := v.(type) {
	case string:
		r = strings.NewReader(v)
	case []byte:
		r = bytes.NewReader(v)
	default:
		panic(pkgName + ": jsonParse only accepts string or []byte: " + reflect.TypeOf(v).String())
	}
	var out interface{}
	jd := json.NewDecoder(r)
	if err := jd.Decode(&out); err != nil {
		panic(pkgName + ": unable to decode from json: " + err.Error())
	}
	return out
}

func init() {
	Register("jsonStringify", jsonStringify)
	Register("jsonPretty", jsonPretty)
	Register("jsonParse", jsonParse)
}
