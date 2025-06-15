//go:build stxOcaml || stxAll

package lexers

func init() {
	syntaxMap["ocaml"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: ocaml
rules:
    - constant.number: "-?0[bB][01][01_]*"
    - constant.number: "-?0[oO][0-7][0-7_]*"
    - constant.number: "-?\\d[\\d_]*"
    - constant.number: "-?0[xX][0-9a-fA-F][0-9a-fA-F_]*"
    - constant.number: "-?\\d[\\d_]*.\\d[\\d_]*([eE][+-]\\d[\\d_]*.\\d[\\d_]*)?"
    - constant.number: "-?0[xX][0-9a-fA-F][0-9a-fA-F_]*.[0-9a-fA-F][0-9a-fA-F_]*([pP][+-][0-9a-fA-F][0-9a-fA-F_]*.[0-9a-fA-F][0-9a-fA-F_]*)?"
    - comment:
        start: "\\(\\*"
        end: "\\*\\)"
        rules: []`)
	}}
}
