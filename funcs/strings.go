package funcs

import (
	"strings"
)

func init() {
	Register("hasPrefix", strings.HasPrefix)
	Register("hasSuffix", strings.HasSuffix)
	Register("contains", strings.Contains)
	Register("toLower", strings.ToLower)
	Register("toUpper", strings.ToUpper)
	Register("toTitle", strings.ToTitle)
	Register("repeat", strings.Repeat)
	Register("join", strings.Join)
	Register("split", strings.Split)
	Register("trim", strings.Trim)
	Register("trimLeft", strings.TrimLeft)
	Register("trimRight", strings.TrimRight)
	Register("trimSpace", strings.TrimSpace)
	Register("trimPrefix", strings.TrimPrefix)
	Register("trimSuffix", strings.TrimSuffix)
}
