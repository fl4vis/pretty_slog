package pretty

import (
	"fmt"
)

const (
	RESET = "\033[0m"

	BLACK  = "40"
	RED    = "31"
	GREEN  = "32"
	YELLOW = "33"
	BLUE   = "34"
	PURPLE = "35"
	CYAN   = "36"
	WHITE  = "37"
)

func Colorize(colorCode string, text string) string {
	return fmt.Sprintf("\033[%sm%s%s", colorCode, text, RESET)
}
