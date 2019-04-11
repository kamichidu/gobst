package funcs

import (
	"github.com/stoewer/go-strcase"
)

func init() {
	Register("snakeCase", strcase.SnakeCase)
	Register("lowerCamelCase", strcase.LowerCamelCase)
	Register("upperCamelCase", strcase.UpperCamelCase)
	Register("kebabCase", strcase.KebabCase)
}
