//go:build stxDiff || stxAll

package lexers

func init() {
	syntaxMap["diff"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: diff
rules:
    - statement: "(^\\+\\+\\+.*)"
    - statement: "(^---.*)"
    - type: "(^@@.*)"
    - constant.number: "(^\\+.*)"
    - preproc: "(^-.*)"`)
	}}
}
