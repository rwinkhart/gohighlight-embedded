//go:build stxDiff || stxAll

package syntax

func init() {
	syntaxMap["diff"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: diff

detect:
    filename: "\\.diff$"
    header: "^((---.*)|(\\+\\+\\+.*))"

rules:
    - statement: "(^\\+\\+\\+.*)"
    - statement: "(^---.*)"
    - type: "(^@@.*)"
    - constant.number: "(^\\+.*)"
    - preproc: "(^-.*)"`)
	}}
}
