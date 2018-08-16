package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

type flagVariable map[string]interface{}

func (v *flagVariable) Set(s string) error {
	if !strings.Contains(s, "=") {
		return errors.New("value must be `name=value` form")
	}
	if *v == nil {
		*v = flagVariable{}
	}
	eles := strings.SplitN(s, "=", 2)
	if len(eles) != 2 {
		panic("len(eles) is not 2 length")
	}
	(*v)[eles[0]] = eles[1]
	return nil
}

func (v flagVariable) String() string {
	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var eles []string
	for _, k := range keys {
		eles = append(eles, k+"="+fmt.Sprint(v[k]))
	}
	return strings.Join(eles, ";")
}

type flagVariableFile map[string]interface{}

func (v *flagVariableFile) Set(s string) error {
	file, err := os.Open(s)
	if err != nil {
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	return dec.Decode(v)
}

func (v flagVariableFile) String() string {
	return flagVariable(v).String()
}

type options struct {
	Variable flagVariable

	OutFile string

	Dir string

	Pat string

	ShowVersion bool
}

func makeTemplate(name string, r io.Reader, dir string, pat *regexp.Regexp) (*template.Template, error) {
	text, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tpl := template.New("")
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() {
			return nil
		}
		// info.Name() is a basename
		if !pat.MatchString(info.Name()) {
			return nil
		}
		// template name is slash separated relative path
		name := filepath.ToSlash(path)
		abspath := filepath.Join(dir, path)
		text, err := ioutil.ReadFile(abspath)
		if err != nil {
			return err
		}
		// set outside variable
		// want to use relpath as a template name, use Parse() not ParseFiles()
		tpl, err = tpl.New(name).Parse(string(text))
		return err
	})
	if err != nil {
		return nil, err
	}
	return tpl.New(name).Parse(string(text))
}

func run(stdin io.Reader, stdout, stderr io.Writer, args []string) int {
	var opts options
	flg := flag.NewFlagSet(filepath.Base(args[0]), flag.ContinueOnError)
	flg.SetOutput(stderr)
	flg.Usage = func() {
		fmt.Fprint(flg.Output(), "Synopsis:\n  execute go text/template via cli\n\n")
		fmt.Fprintf(flg.Output(), "Usage:\n  %s [options] {template file}\n\n", flg.Name())
		fmt.Fprint(flg.Output(), "Options:\n")
		flg.PrintDefaults()
	}
	flg.Var(&opts.Variable, "var", "template variable(s) with `NAME=VALUE` form")
	flg.Var((*flagVariableFile)(&opts.Variable), "var-file", "template variable(s) from given json `FILE`")
	flg.Bool("h", false, "show help message")
	flg.BoolVar(&opts.ShowVersion, "v", false, "show version")
	flg.StringVar(&opts.OutFile, "o", "-", "output `FILE`")
	flg.StringVar(&opts.Dir, "dir", ".", "the `DIRECTORY` that finding template files matching glob")
	flg.StringVar(&opts.Pat, "pat", `\.tmpl$`, "template file `REGEXP`")
	if err := flg.Parse(args[1:]); err == flag.ErrHelp {
		return 0
	} else if err != nil {
		fmt.Println(stderr, err)
		return 128
	} else if flg.NArg() != 1 {
		flg.Usage()
		return 128
	}

	filterRe, err := regexp.Compile(opts.Pat)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 128
	}

	var r io.Reader
	if filename := flg.Arg(0); filename == "-" {
		r = stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		r = file
	}
	if closer, ok := r.(io.Closer); ok {
		defer closer.Close()
	}

	var w io.Writer
	if opts.OutFile == "-" {
		w = stdout
	} else {
		file, err := os.Create(opts.OutFile)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		w = file
	}
	if closer, ok := w.(io.Closer); ok {
		defer closer.Close()
	}

	tpl, err := makeTemplate(flg.Name(), r, opts.Dir, filterRe)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if err := tpl.Execute(w, (map[string]interface{})(opts.Variable)); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Stdin, os.Stdout, os.Stderr, os.Args))
}
