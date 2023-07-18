# Highlight
[![Go Report Card](https://goreportcard.com/badge/github.com/jessp01/gohighlight)](https://goreportcard.com/report/github.com/jessp01/gohighlight)
[![GoDoc](https://godoc.org/github.com/jessp01/gohighlight?status.svg)](http://godoc.org/github.com/jessp01/gohighlight)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jessp01/gohighlight/blob/master/LICENSE)

`gohighlight` is a syntax highlighter of programming languages, config formats and UNIX commands. 
It allows you to pass in a string and get back all the information you need to properly highlight it..
This repo includes over a 100 different lexers and contributions are most welcome:)

See [Revising and adding new lexers](#revising-and-adding-new-lexers) for more info.

## A note about this repo

This is a fork of [zyedidia/highlight](https://github.com/zyedidia/highlight). 
I originally submitted pulls against it but it seems to be unmaintained (last commit was in 2020).

Below is a recap of the main changes made since:
- Helper method to parse syntax files
- Better exception handling
- Corrections/additions to existing lexers
- Addition of new lexers
- Revisions to the code examples
- Basic unit tests (still more to do on that)

## Installation

```
$ go get github.com/jessp01/highlight
```

Be sure to point your code to the correct path of `syntax_files`.

## Basic Usage

Below is a simple example for highlighting a string (a Go snippet in this case). 
It uses `github.com/fatih/color` to actually colorize the output to the console.

**NOTE: A more comprehensive example is [zaje](https://github.com/jessp01/zaje); a Syntax highlighter to cover all your shell needs (it can replace `cat` and `tail`).**

```go
package main

import (
    "fmt"
    "io/ioutil"
    "strings"

    "github.com/fatih/color"
    "github.com/jessp01/highlight"
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
    syntaxFile, lerr := ioutil.ReadFile("highlight/syntax_files/go.yaml")
    if lerr != nil {
        fmt.Println(lerr)
	return
    }    

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

If you would like to automatically detect the filetype based on the filename, you can use the `DetectFiletype()` function:

```go
// Name of the file
filename := ...
// The first line of the file (needed to check the filetype by header: e.g. `#!/bin/bash` means shell)
firstLine := ...

// Parse all the syntax files in an array with type []*highlight.Def
var defs []*highlight.Def
...

def := highlight.DetectFiletype(defs, filename, firstLine)
fmt.Println("Filetype is", def.FileType)
```

## Revising and adding new lexers

Lexers are YML files that live under the [syntax\_files](./syntax_files) dir.

They can be loaded individually:
```go
    syntaxFile, lerr := ioutil.ReadFile("highlight/syntax_files/go.yaml")
    syntaxDef, err := highlight.ParseDef(syntaxFile)
```

Or, you can scan the dir and load them all using the `highlight.ParseSyntaxFiles(syn_dir, &defs)()` helper function.
