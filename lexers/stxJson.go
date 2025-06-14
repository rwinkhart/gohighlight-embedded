//go:build stxJson || stxAll

package lexers

func init() {
	syntaxMap["json"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: json
rules:
    - constant.number: "\\b[-+]?([1-9][0-9]*|0[0-7]*|0x[0-9a-fA-F]+)([uU][lL]?|[lL][uU]?)?\\b"
    - constant.number: "\\b[-+]?([0-9]+\\.[0-9]*|[0-9]*\\.[0-9]+)([EePp][+-]?[0-9]+)?[fFlL]?"
    - constant.number: "\\b[-+]?([0-9]+[EePp][+-]?[0-9]+)[fFlL]?"
    - constant: "\\b(null)\\b"
    - constant: "\\b(true|false)\\b"
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
            - constant.specialChar: "\\\\."
    - statement: "\\\"(\\\\\"|[^\"])*\\\"[[:space:]]*:\"  \"'(\\'|[^'])*'[[:space:]]*:"
    - constant: "\\\\u[0-9a-fA-F]{4}|\\\\[bfnrt'\"/\\\\]"`)
	}}
}
