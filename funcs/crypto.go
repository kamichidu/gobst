package funcs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func cryptoMd5(v string) string {
	b := md5.Sum([]byte(v))
	return hex.EncodeToString(b[:])
}

func cryptoMd5file(v string) string {
	data, err := os.ReadFile(v)
	if err != nil {
		panic(fmt.Sprintf("%v: unable to opena file: %v: %v", pkgName, v, err))
	}
	b := md5.Sum(data)
	return hex.EncodeToString(b[:])
}

func init() {
	Register("md5", cryptoMd5)
	Register("md5file", cryptoMd5file)
}
