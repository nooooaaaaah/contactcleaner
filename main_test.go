package contactcleaner

import (
	"strings"
	"testing"
)

// parseTelephone is a mock function for demonstration.
// Replace this with your actual implementation.
func parseTelephone(line string) Telephone {
	var tel Telephone
	parts := strings.Split(line, ":")
	if len(parts) == 2 {
		tel.Number = parts[1]
		typePart := strings.Split(parts[0], ";")
		if len(typePart) > 1 {
			tel.Type = typePart[1:]
		}
	}
	return tel
}

func TestParseTelephone(t *testing.T) {
	tests := []struct {
		line   string
		number string
		typ    []string
	}{
		{"TEL:123-456-7890", "123-456-7890", []string{""}},
		{"TEL;TYPE=WORK:555-123-4567", "555-123-4567", []string{"TYPE=WORK"}},
		{"TEL;TYPE=CELL:987-654-3210", "987-654-3210", []string{"TYPE=CELL"}},
		{"TEL;TYPE=HOME;PREF:+44 20 7946 0200", "+44 20 7946 0200", []string{"TYPE=HOME;PREF"}},
		{"TEL;bogus=data:123-456-7890", "123-456-7890", []string{"bogus=data"}},
	}

	for _, test := range tests {
		telephone := parseTelephone(test.line)
		if telephone.Number != test.number {
			t.Errorf("Expected phone number '%s', got '%s'", test.number, telephone.Number)
		}
		if telephone.Type[0] != test.typ[0] {
			t.Errorf("Expected phone type '%s', got '%s'", test.typ, telephone.Type)
		}
	}
}
