//go:build stxMan || stxAll

package syntax

func init() {
	syntaxMap["man"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: man
rules:
    - green: "\\.(S|T)H.*$"
    - brightgreen: "\\.(S|T)H|\\.TP"
    - brightred: "\\.(BR?|I[PR]?).*$"
    - brightblue: "\\.(BR?|I[PR]?|PP)"
    - brightwhite: "\\\\f[BIPR]"
    - yellow: "\\.(br|DS|RS|RE|PD)"`)
	}}
}
