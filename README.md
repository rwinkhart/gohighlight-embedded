# Highlight

`gohighlight-embedded` is a syntax highlighter of programming languages and config formats.
It allows you to pass in a string and get back all the information you need to properly highlight it.
This repo includes lexers for many common languages.

## A note about this repo

This is a fork of [jessp01/gohighlight](https://github.com/jessp01/gohighlight).

It is meant to be a more minimal version with the syntax files embedded in the Go code (no external .yaml files to load).

## Installation

```sh
$ go get github.com/rwinkhart/gohighlight-embedded
```

Be sure to point your code to the correct path of `syntax_files`.

## Basic Usage

Below is a simple example for highlighting a string (a Go snippet in this case).
It uses `github.com/fatih/color` to actually colorize the output to the console.

```go
package main

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
    "github.com/rwinkhart/gohighlight-embedded"
    "github.com/rwinkhart/gohighlight-embedded/syntax"
)

func main() {
    // Here is the go code we will highlight
    inputString := `package main

import "fmt"

// A hello world program
func helloWorld() {
    fmt.Println("Hello world")
}`

    // Load the go syntax file
    // Make sure that the syntax_files directory is in the current directory
    syntaxFile := syntax.GetGo()

    // Parse it into a `*highlight.Def`
    syntaxDef, err := highlight.ParseDef(syntaxFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Make a new highlighter from the definition
    h := highlight.NewHighlighter(syntaxDef)
    // Highlight the string
    // Matches is an array of maps which point to groups
    // matches[lineNum][colNum] will give you the change in group at that line and column number
    // Note that there is only a group at a line and column number if the syntax highlighting changed at that position
    matches := h.HighlightString(inputString)

    // We split the string into a bunch of lines
    // Now we will print the string
    lines := strings.Split(inputString, "\n")
    for lineN, l := range lines {
	    colN := 0
	    for _, c := range l {
		    if group, ok := matches[lineN][colN]; ok {
			    switch group {
			    case highlight.Groups["statement"]:
				    fallthrough
			    case highlight.Groups["green"]:
				    color.Set(color.FgGreen)

			    case highlight.Groups["identifier"]:
				    fallthrough
			    case highlight.Groups["blue"]:
				    color.Set(color.FgHiBlue)

			    case highlight.Groups["preproc"]:
				    //fallthrough
				    //case highlight.Groups["high.red"]:
				    color.Set(color.FgHiRed)

			    case highlight.Groups["special"]:
				    fallthrough
			    case highlight.Groups["red"]:
				    color.Set(color.FgRed)

			    case highlight.Groups["constant.string"]:
				    fallthrough
			    case highlight.Groups["constant"]:
				    fallthrough
			    case highlight.Groups["constant.number"]:
				    fallthrough
			    case highlight.Groups["cyan"]:
				    color.Set(color.FgCyan)

			    case highlight.Groups["constant.specialChar"]:
				    fallthrough
			    case highlight.Groups["magenta"]:
				    color.Set(color.FgHiMagenta)

			    case highlight.Groups["type"]:
				    fallthrough
			    case highlight.Groups["yellow"]:
				    color.Set(color.FgYellow)

			    case highlight.Groups["comment"]:
				    fallthrough
			    case highlight.Groups["high.green"]:
				    color.Set(color.FgHiGreen)
			    default:
				    color.Unset()
			    }
		    }
		    fmt.Print(string(c))
		    colN++
	    }
	    if group, ok := matches[lineN][colN]; ok {
		    if group == highlight.Groups["default"] || group == highlight.Groups[""] {
			    color.Unset()
		    }
	    }

	    color.Unset()
	    fmt.Print("\n")
    }
}
```
