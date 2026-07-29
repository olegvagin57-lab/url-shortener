package generator

import (
	"strings"
)

var alphabet string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateCode(id int) string {
	var finalStr strings.Builder
	for id > 0 {
		idx := id % 62
		finalStr.WriteByte(alphabet[idx])
		id = id / 62
	}

	result := finalStr.String()

	return result
}
