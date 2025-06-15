//go:build stxLisp || stxAll

package lexers

func init() {
	syntaxMap["lisp"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: lisp
rules:
    - default: "\\([a-z-]+"
    - symbol: "\\(([\\-+*/<>]|<=|>=)|'"
    - constant.number: "\\b[0-9]+b>"
    - special: "\\bnil\\b"
    - preproc: "\\b[tT]b>"
    - constant.string: "\\\"(\\\\.|[^\"])*\\\""
    - constant.specialChar: "'[A-Za-z][A-Za-z0-9_-]+"
    - constant.specialChar: "\\\\.?"
    - comment: "(^|[[:space:]]);.*"
    - indent-char.whitespace: "[[:space:]]+$"
    - indent-char: "	+ +| +	+"`)
	}}
}
