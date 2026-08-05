package generator

import "testing"

func TestGenerator_GenerateCode(t *testing.T) {
	g := NewGenerator()

	code := g.GenerateCode()

	if code == "" {
		t.Error("Generated empty string")
	}

	seen := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		code := g.GenerateCode()
		if seen[code] {
			t.Errorf("Duplicate code generated: %s", code)
		}
		seen[code] = true
	}
}
