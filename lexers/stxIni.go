//go:build stxIni || stxAll

package lexers

func init() {
	syntaxMap["ini"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: ini
rules:
    - constant.bool: "\\b(true|false)\\b"
    - identifier: "^[[:space:]]*[^=]*="
    - special: "^[[:space:]]*\\[.*\\]$"
    - statement: "[=;]"
    - comment: "(^|[[:space:]])#([^{].*)?$"
    - constant.string: "\"(\\\\.|[^\"])*\"|'(\\\\.|[^'])*'"`)
	}}
}
