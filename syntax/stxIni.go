//go:build stxIni || stxAll

package syntax

func init() {
	syntaxMap["ini"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: ini
rules:
    - constant.bool.true: "\\btrue\\b"
    - constant.bool.false: "\\bfalse\\b"
    - identifier: "^[[:space:]]*[^=]*="
    - special: "^[[:space:]]*\\[.*\\]$"
    - statement: "[=;]"
    - comment: "(^|[[:space:]])#([^{].*)?$"
    - constant.string: "\"(\\\\.|[^\"])*\"|'(\\\\.|[^'])*'"`)
	}}
}
