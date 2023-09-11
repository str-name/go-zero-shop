package tool

import "strings"

func GetFirstImg(str string) string {
	return strings.Split(str, ",")[0]
}
