gobst
========================================================================================================================
Golang text/template for cli.

Installation
------------------------------------------------------------------------------------------------------------------------
Visit [releases](https://github.com/kamichidu/gobst/releases).

Or

```
go get github.com/kamichidu/gobst
```

Template Functions
------------------------------------------------------------------------------------------------------------------------
Standard golang text/template functions are listed [here](https://golang.org/pkg/text/template/#hdr-Functions).

gobst defined functions are listed below:

- snakeCase(s string) string
- lowerCamelCase(s string) string
- upperCamelCase(s string) string
- kebabCase(s string) string
- hasPrefix(s, prefix string) bool
- hasSuffix(s, suffix string) bool
- contains(s, substr string) bool
- toLower(s string) string
- toUpper(s string) string
- toTitle(s string) string
- repeat(s string, n int) string
- join(l []string, sep string) string
- split(s, sep string) string
- trim(s, cutset string) string
- trimLeft(s, cutset string) string
- trimRight(s, cutset string) string
- trimSpace(s string) string
- trimPrefix(s, prefix string) string
- trimSuffix(s, suffix string) string
- jsonStringify(v interface{}) string
- jsonPretty(v interface{}) string
- jsonParse(data string) interface{}
- jsonParse(data []byte) interface{}
- slice(args ...interface{}) []interface{}
- keys(v map[string]interface{}) []string
- command(args ...interface{}) map[string]interface{}
