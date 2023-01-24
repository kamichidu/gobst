package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("stdin template", func(t *testing.T) {
		var stdin bytes.Buffer
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		_, err := stdin.WriteString("hi {{ .name }}!")
		if err != nil {
			panic(err)
		}
		code := Run(&stdin, &stdout, &stderr, []string{
			"gobst",
			"-var", "name=kamichidu",
			"-",
		})
		assert.Equal(t, 0, code)
		assert.Equal(t, "hi kamichidu!", stdout.String())
		assert.Equal(t, "", stderr.String())
	})
	t.Run("flat", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		code := Run(nil, &stdout, &stderr, []string{
			"gobst",
			"-var", "name=kamichidu",
			"./testdata/flat/simple.tmpl",
		})
		assert.Equal(t, 0, code)
		assert.Equal(t, "hi kamichidu!", stdout.String())
		assert.Equal(t, "", stderr.String())
	})
	t.Run("nested", func(t *testing.T) {
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		code := Run(nil, &stdout, &stderr, []string{
			"gobst",
			"-var", "name=kamichidu",
			"-dir", "./testdata/nested/templates/",
			"./testdata/nested/root.tmpl",
		})
		assert.Equal(t, 0, code)
		assert.Equal(t, `hi kamichidu!
this is nested template content.
the name is "kamichidu".
`, stdout.String())
		assert.Equal(t, "", stderr.String())
	})
}
