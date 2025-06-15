//go:build stxC || stxAll

package lexers

func init() {
	syntaxMap["c"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: c
rules:
    - identifier: "\\b[A-Z_][0-9A-Z_]+\\b"
    - statement: "([a-zA-Z][a-zA-Z0-9_]*)[[:space:]]*\\("
    - type: "\\b(float|double|char|bool|int|short|long|sizeof|enum|void|static|const|struct|union|typedef|extern|(un)?signed|inline)\\b"
    - type: "\\b((s?size)|((u_?)?int(8|16|32|64|ptr)))_t\\b"
    - statement: "\\b(typename|mutable|volatile|register|explicit)\\b"
    - statement: "\\b(for|if|while|do|else|case|default|switch)\\b"
    - statement: "\\b(try|throw|catch|operator|new|delete)\\b"
    - statement: "\\b(goto|continue|break|return)\\b"
    - preproc: "^[[:space:]]*#[[:space:]]*(define|pragma|include|(un|ifn?)def|endif|el(if|se)|if|warning|error)"
    - constant: "'([^'\\\\]|(\\\\[\"'abfnrtv\\\\]))'"
    - constant: "'\\\\(([0-3]?[0-7]{1,2}))'"
    - constant: "'\\\\x[0-9A-Fa-f]{1,2}'"
    - statement: "__attribute__[[:space:]]*\\(\\([^)]*\\)\\)"
    - statement: "__(aligned|asm|builtin|hidden|inline|packed|restrict|section|typeof|weak)__"
    - symbol.operator: "([.:;,+*|=!\\%]|<|>|/|-|&)"
    - symbol.brackets: "[(){}]|\\[|\\]"
    - constant.number: "(\\b[0-9]+\\b|\\b0x[0-9A-Fa-f]+\\b)"
    - constant.number: "NULL"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - preproc: "..+"
            - constant.specialChar: "\\\\."
    - comment:
        start: "//"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
