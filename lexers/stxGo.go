//go:build stxGo || stxAll

package lexers

func init() {
	syntaxMap["go"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: go
rules:
    - special: "\\b(break|case|continue|default|go|goto|range|return)\\b"
    - statement: "\\b(else|for|if|switch)\\b"
    - preproc: "\\b(package|import|const|var|type|struct|func|go|defer|iota)\\b"
    - symbol.operator: "[-+/*=<>!~%&|^]|:="
    - special: "[a-zA-Z0-9]*\\("
    - symbol: "(,|\\.)"
    - type: "\\b(u?int(8|16|32|64)?|float(32|64)|complex(64|128))\\b"
    - type: "\\b(uintptr|byte|rune|string|interface|bool|map|chan|error)\\b"
    - keyword: "\\b(struct)\\b"
    - constant.bool: "\\b(true|false|nil)\\b"
    - symbol.brackets: "(\\{|\\})"
    - symbol.brackets: "(\\(|\\))"
    - symbol.brackets: "(\\[|\\])"
    - constant.number: "\\b([0-9]+|0x[0-9a-fA-F]*)\\b|'.'"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "%."
            - constant.specialChar: "\\\\[abfnrtv'\\\"\\\\]"
            - constant.specialChar: "\\\\([0-7]{3}|x[A-Fa-f0-9]{2}|u[A-Fa-f0-9]{4}|U[A-Fa-f0-9]{8})"
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - error: "..+"
            - constant.specialChar: "%."
            - constant.specialChar: "\\\\[abfnrtv'\\\"\\\\]"
            - constant.specialChar: "\\\\([0-7]{3}|x[A-Fa-f0-9]{2}|u[A-Fa-f0-9]{4}|U[A-Fa-f0-9]{8})"
    - constant.string:
        start: "` + "`" + `"
        end: "` + "`" + `"
        rules: []
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
