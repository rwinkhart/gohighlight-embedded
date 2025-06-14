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
Below is a simple example for highlighting a Go snippet with ANSI colors.

It must be built with `-tags=stxGo` or `-tags=stxAll` to include the Go syntax layer.

For a more advanced example on manually colorizing the output or colorizing it in a
different, non-ANSI format, see the `ansi.Colorize` function.

```go
package main

import (
	"fmt"
	"os"

	"github.com/rwinkhart/go-highlite/ansi"
)

func main() {
	// Here is the go code we will highlight
	inputString := `package main

import "fmt"

// A hello world program
func helloWorld() {
    fmt.Println("Hello world")
}`

	result, err := ansi.Colorize(inputString, "go")
	if err != nil {
		fmt.Println("Failed to colorize input: " + err.Error())
		os.Exit(1)
	}
	fmt.Println(result)
}

```
