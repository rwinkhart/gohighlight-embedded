package ansi

import (
	"errors"
	"strings"

	"github.com/rwinkhart/go-highlite"
	"github.com/rwinkhart/go-highlite/syntax"
)

// Colorize takes a string of code and a language ID, and returns the code
// with ANSI color codes applied for syntax highlighting.
func Colorize(inputCode, languageID string) (string, error) {
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
				var colorCode string
				switch group {
				case highlite.Groups["statement"]:
					colorCode = "\033[32m" // Green
				case highlite.Groups["identifier"]:
					colorCode = "\033[94m" // Bright Blue
				case highlite.Groups["preproc"]:
					colorCode = "\033[91m" // Bright Red
				case highlite.Groups["special"], highlite.Groups["red"]:
					colorCode = "\033[31m" // Red
				case highlite.Groups["constant.string"], highlite.Groups["constant"],
					highlite.Groups["constant.number"], highlite.Groups["cyan"]:
					colorCode = "\033[36m" // Cyan
				case highlite.Groups["constant.specialChar"], highlite.Groups["magenta"]:
					colorCode = "\033[95m" // Bright Magenta
				case highlite.Groups["type"]:
					colorCode = "\033[33m" // Yellow
				case highlite.Groups["comment"], highlite.Groups["brightgreen"]:
					colorCode = "\033[92m" // Bright Green
				default:
					colorCode = "\033[0m" // Reset
				}
				outputString.WriteString(colorCode)
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
