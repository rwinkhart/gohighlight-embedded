//go:build stxCpp || stxAll

package lexers

func init() {
	syntaxMap["cpp"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: c++
rules:
    - identifier: "\\b[A-Z_][0-9A-Z_]+\\b"
    - type: "\\b(auto|float|double|bool|char|int|short|long|sizeof|enum|void|static|const|constexpr|struct|union|typedef|extern|(un)?signed|inline)\\b"
    - type: "\\b((s?size)|((u_?)?int(8|16|32|64|ptr)))_t\\b"
    - statement: "\\b(class|namespace|template|public|protected|private|typename|this|friend|virtual|using|mutable|volatile|register|explicit)\\b"
    - statement: "\\b(for|if|while|do|else|case|default|switch)\\b"
    - statement: "\\b(try|throw|catch|operator|new|delete)\\b"
    - statement: "\\b(goto|continue|break|return)\\b"
    - preproc: "^[[:space:]]*#[[:space:]]*(define|pragma|include|(un|ifn?)def|endif|el(if|se)|if|warning|error)"
    - constant: "('([^'\\\\]|(\\\\[\"'abfnrtv\\\\]))'|'\\\\(([0-3]?[0-7]{1,2}))'|'\\\\x[0-9A-Fa-f]{1,2}')"
    - statement: "(__attribute__[[:space:]]*\\(\\([^)]*\\)\\)|__(aligned|asm|builtin|hidden|inline|packed|restrict|section|typeof|weak)__)"
    - symbol.operator: "([.:;,+*|=!\\%]|<|>|/|-|&)"
    - symbol.brackets: "[(){}]|\\[|\\]"
    - constant.number: "(\\b[0-9]+\\b|\\b0x[0-9A-Fa-f]+\\b)"
    - constant.bool: "(\\b(true|false)\\b|NULL)"
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
