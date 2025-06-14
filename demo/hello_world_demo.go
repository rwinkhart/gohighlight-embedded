package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/rwinkhart/go-highlite/ansi"
)

func main() {
	// Slice of strings containing Hello World examples for different languages
	helloWorldExamples := []struct {
		language string
		code     string
	}{
		{
			language: "go",
			code: `// Hello World in Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`,
		},
		{
			language: "python",
			code: `# Hello World in Python
print("Hello, World!")`,
		},
		{
			language: "java",
			code: `// Hello World in Java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}`,
		},
		{
			language: "c",
			code: `// Hello World in C
#include <stdio.h>

int main() {
    printf("Hello, World!\n");
    return 0;
}`,
		},
		{
			language: "cpp",
			code: `// Hello World in C++
#include <iostream>

int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}`,
		},
		{
			language: "js",
			code: `// Hello World in JavaScript
console.log("Hello, World!");`,
		},
		{
			language: "rust",
			code: `// Hello World in Rust
fn main() {
    println!("Hello, World!");
}`,
		},
		{
			language: "ruby",
			code: `# Hello World in Ruby
puts "Hello, World!"`,
		},
		{
			language: "php",
			code: `<?php
// Hello World in PHP
echo "Hello, World!\n";
?>`,
		},
		{
			language: "swift",
			code: `// Hello World in Swift
print("Hello, World!")`,
		},
		{
			language: "scala",
			code: `// Hello World in Scala
object HelloWorld {
  def main(args: Array[String]): Unit = {
    println("Hello, World!")
  }
}`,
		},
		{
			language: "haskell",
			code: `-- Hello World in Haskell
main :: IO ()
main = putStrLn "Hello, World!"`,
		},
		{
			language: "perl",
			code: `# Hello World in Perl
print "Hello, World!\n";`,
		},
		{
			language: "lua",
			code: `-- Hello World in Lua
print("Hello, World!")`,
		},
		{
			language: "shell",
			code: `#!/bin/bash
# Hello World in Shell
echo "Hello, World!"`,
		},
	}

	fmt.Print("=== Go-Highlite Demo: Hello World Examples ===\n\n")

	for _, example := range helloWorldExamples {
		fmt.Printf("Language: %s\n\n", example.language)

		colorized, err := ansi.Colorize(example.code, example.language)
		if err != nil {
			log.Printf("Error colorizing %s: %v\n", example.language, err)
			fmt.Println("(Unable to highlight this code)")
		} else {
			fmt.Println(colorized)
		}

		fmt.Println("\n" + strings.Repeat("-", 50) + "\n")
	}
}
