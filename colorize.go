package pretty_slog

import (
	"bytes"
)

const (
	RESET  = "\x1b[0m"
	RED    = "\x1b[31m"
	GREEN  = "\x1b[32m"
	YELLOW = "\x1b[33m"
	BLUE   = "\x1b[34m"
	PURPLE = "\x1b[35m"
	CYAN   = "\x1b[36m"
	GRAY   = "\x1b[90m"
	WHITE  = "\x1b[37m"
)

const (
	LEFT = iota
	RIGHT
)

func Colorize(color, text string) string {
	return color + text + RESET
}

func ColorizeJSON(data []byte) string {
	var buf bytes.Buffer

	// \x1b[0m
	resetSequence := []byte{27, 91, 48, 109}
	needsCheck := false
	isBracket := false
	isQuote := false
	semicolonPosition := LEFT

	for i := range data {
		char := data[i]

		switch char {
		case '{', '}', '[', ']', '"', ',':
			buf.WriteString(WHITE)
			buf.WriteByte(char)
			buf.WriteString(RESET)

			needsCheck = true
			isQuote = false

			if char == '{' || char == '}' {
				semicolonPosition = LEFT
				isBracket = false
			} else if char == ',' {
				semicolonPosition = LEFT
			} else if char == '[' {
				isBracket = true
			} else if char == ']' {
				isBracket = false
			} else if char == '"' {
				isQuote = true
			}

		case ':':
			buf.WriteString(WHITE)
			buf.WriteByte(char)
			buf.WriteString(RESET)

			if semicolonPosition == 0 {
				semicolonPosition = RIGHT
			} else {
				semicolonPosition = LEFT
			}

		default:
			if needsCheck {
				bufLen := buf.Len()
				if bufLen >= 4 {
					last4 := buf.Bytes()[bufLen-4:]
					// Check for reset sequence -> \x1b[0m
					if bytes.Equal(last4, resetSequence) {
						if isBracket {
							semicolonPosition = RIGHT
						}

						// Determine color based on value type
						color := PURPLE
						if semicolonPosition == LEFT {
							color = BLUE
						} else if semicolonPosition == RIGHT {
							if isQuote {
								color = GREEN
							}
						}
						buf.WriteString(color)
					}
				}
			}

			buf.WriteByte(data[i])

		}
	}

	return buf.String()
}
