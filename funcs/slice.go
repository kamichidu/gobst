package funcs

func slice(args ...interface{}) (l []interface{}) {
	for _, a := range args {
		l = append(l, a)
	}
	return l
}

func init() {
	Register("slice", slice)
}
