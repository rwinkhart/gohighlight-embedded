package ansi

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rwinkhart/go-highlite"
	"github.com/rwinkhart/go-highlite/syntax"
)

// Colorize takes a string of code and a language ID, and returns the code
// with ANSI color codes applied for syntax highlighting.
func Colorize(inputCode, languageID string, startingColor uint8) (string, error) {
	// Load and parse the go syntax layer into a `*highlight.Def`
	syntaxDef, err := highlite.ParseDef(syntax.Get(languageID))
	if err != nil {
		return "", errors.New("unable to parse syntax layer: " + err.Error())
	}

	// Make a new highlighter from the definition
	h := highlite.NewHighlighter(syntaxDef)

	// Highlight the string
	matches := h.HighlightString(inputCode)

	// We split the string into a bunch of lines
	lines := strings.Split(inputCode, "\n")
	var outputString strings.Builder

	for lineN, l := range lines {
		colN := 0
		for _, c := range l {
			if group, ok := matches[lineN][colN]; ok {
				var colorOffset uint8
				switch group {
				case highlite.Groups["comment"]: // TODO static color
					colorOffset = 0
				case highlite.Groups["constant"]:
					colorOffset = 2
				case highlite.Groups["constant.bool"]:
					colorOffset = 4
				case highlite.Groups["constant.number"]:
					colorOffset = 6
				case highlite.Groups["constant.specialChar"]:
					colorOffset = 8
				case highlite.Groups["constant.string"]:
					colorOffset = 10
				case highlite.Groups["error"]: // TODO static color
					colorOffset = 12
				case highlite.Groups["identifier"]:
					colorOffset = 14
				case highlite.Groups["identifier.macro"]:
					colorOffset = 16
				case highlite.Groups["identifier.var"]:
					colorOffset = 18
				case highlite.Groups["keyword"]:
					colorOffset = 20
				case highlite.Groups["preproc"]:
					colorOffset = 22
				case highlite.Groups["special"]:
					colorOffset = 24
				case highlite.Groups["statement"]:
					colorOffset = 26
				case highlite.Groups["symbol"]:
					colorOffset = 28
				case highlite.Groups["symbol.brackets"]:
					colorOffset = 30
				case highlite.Groups["symbol.operator"]:
					colorOffset = 32
				case highlite.Groups["todo"]:
					colorOffset = 34
				case highlite.Groups["type"]:
					colorOffset = 36
				default:
					colorOffset = 255
				}
				if colorOffset == 255 {
					outputString.WriteString("\033[0m") // Reset/default color
				} else {
					outputString.WriteString(fmt.Sprintf("\033[38;5;%dm", startingColor+colorOffset))
				}
			}
			outputString.WriteRune(c)
			colN++
		}
		outputString.WriteString("\033[0m") // Reset colors at end of line
		if lineN < len(lines)-1 {           // Don't add newline after last line
			outputString.WriteRune('\n')
		}
	}
	return outputString.String(), nil
}
