# Highlite
`go-highlite` is a lightweight ("lite") syntax highlighter of programming languages and config formats.
It allows you to pass in a string and get back all the information you need to properly highlight it.

## A note about this repo
This is an opinionated hard-fork of [jessp01/gohighlight](https://github.com/jessp01/gohighlight).

It is meant to be a more minimal version with the syntax files embedded in the Go code (no external .yaml files to load).

Language support is determined at build time by using build tags. To include a language, you must build with the appropriate tag.
For example, to include Go syntax highlighting, you must build with `-tags=stxGo`.
You may combine as many tags as you like, or use `-tags=stxAll` to include all syntax layers.

### Contributing
If you would like to contribute a new syntax layer, please contribute to the parent repository [jessp01/gohighlight](https://github.com/jessp01/gohighlight).
In an effort to minimize the size of this library, the syntax layers are modified to the point of being much less readable than the original.
Because of this, the parent repository is the canonical source for syntax layers.

I will keep this repository up-to-date with any new syntax layers that are added to the parent repository, provided that they are for non-obsolete
programming languages or config formats (support will not be extended to cover individual shell commands).

Other types of contributions are welcome as long as they do not go against the goal of keeping this library as lightweight as possible.

## Installation
```sh
$ go get github.com/rwinkhart/go-highlite
```

## Basic Usage
Below is a simple example for highlighting a string (a Go snippet in this case).
It uses `github.com/fatih/color` to actually colorize the output to the console.

It must be built with `-tags=stxGo` or `-tags=stxAll` to include the Go syntax layer.

```go
package main

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
    "github.com/rwinkhart/go-highlite"
    "github.com/rwinkhart/go-highlite/syntax"
)

func main() {
    // Here is the go code we will highlight
    inputString := `package main

import "fmt"

// A hello world program
func helloWorld() {
    fmt.Println("Hello world")
}`

    // Load and parse the go syntax layer into a `*highlight.Def`
    syntaxDef, err := highlight.ParseDef(syntax.Get("go"))
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
