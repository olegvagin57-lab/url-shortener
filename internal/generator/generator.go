package generator

import (
	"slices"
	"strings"
)

type Generator struct {
	nextID int
}

var alphabet string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewGenerator() *Generator {
	return &Generator{
		nextID: 100,
	}
}

func (g *Generator) GenerateCode() string {
	var finalStr strings.Builder
	id := g.nextID
	for id > 0 {
		idx := id % 62
		finalStr.WriteByte(alphabet[idx])
		id = id / 62
	}
	result := reverse(finalStr.String())
	g.nextID++
	return result
}

func reverse(s string) string {
	str := []rune(s)
	slices.Reverse(str)
	return string(str)
}
