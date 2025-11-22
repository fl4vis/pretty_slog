package pretty_slog

import (
	"fmt"
	"strings"
	"testing"
)
var JSON = []byte(`{
	  "name": "John Doe",
	  "age": 30,
	  "height": 5.9,
	  "active": true,
	  "verified": false,
	  "balance": null,
	  "email": "john@example.com",
	  "metadata": {
		"created": "2024-01-15",
		"updated": "2024-03-20",
		"tags": ["developer", "golang", "backend"],
		"scores": [95, 87.5, 92, 100],
		"settings": {
		  "notifications": true,
		  "theme": "dark",
		  "language": null
		}
	  },
	  "addresses": [
		{
		  "type": "home",
		  "street": "123 Main St",
		  "city": "Boston",
		  "zip": 12345
		},
		{
		  "type": "work",
		  "street": "456 Tech Blvd",
		  "city": "San Francisco",
		  "zip": 94105
		}
	  ],
	  "preferences": {
		"colors": ["red", "blue", "green"],
		"numbers": [1, 2, 3, 4.5, -10, 0],
		"flags": [true, false, true],
		"mixed": [null, "text", 42, true, false]
	  },
	  "empty_object": {},
	  "empty_array": [],
	  "special_chars": "Hello \"World\" \n\t\\ End",
	  "unicode": "Hello 世界 🌍",
	  "negative": -123.456,
	  "scientific": 1.23e-4,
	  "zero": 0
	}`)

func TestColorizeJson_Visual(t *testing.T) {
	fmt.Println(ColorizeJSON(JSON))
}

func TestColorizeJson_Debug(t *testing.T) {
	result := ColorizeJSON(JSON)
	
	// Print as bytes to see ANSI codes
	fmt.Printf("Raw bytes: %v\n", []byte(result))
	fmt.Printf("As string: %s\n", result)
	fmt.Printf("Length: input=%d output=%d\n", len(JSON), len(result))
	fmt.Printf("%q\n", result)

}

func TestColorizeJson_Text(t *testing.T) {
	text := []byte(`{"name":"value"}`)
	result := ColorizeJSON(text)

	// "name" should have ONE color code before it, not 4 separate ones
	// Bad:  \x1b[36mn\x1b[37ma\x1b[37mm\x1b[37me
	// Good: \x1b[36mname\x1b[0m

	escapeCount := strings.Count(result, "\x1b")

	fmt.Printf("%s\n", result)
	fmt.Printf("%q\n\n", result)
	fmt.Printf("Total \\x1b sequences: %d\n", escapeCount)
	fmt.Printf("Length: input=%d output=%d\n\n", len(text), len(result))

	if escapeCount > 15 {
		t.Errorf("Too many \\x1b! Got %d", escapeCount)
	}
}
