package utils

import (
	"github.com/axgle/mahonia"
)

func GBK2UTF(str string) string {
	return mahonia.NewDecoder("gbk").ConvertString(str)
}

func UTF2GBk(str string) string {
	return mahonia.NewEncoder("gbk").ConvertString(str)
}
