//go:build stxScala || stxAll

package lexers

func init() {
	syntaxMap["scala"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: scala
rules:
    - type: "\\b(boolean|byte|char|double|float|int|long|new|short|this|transient|void)\\b"
    - statement: "\\b(match|val|var|break|case|catch|continue|default|do|else|finally|for|if|return|switch|throw|try|while)\\b"
    - statement: "\\b(def|object|case|trait|lazy|implicit|abstract|class|extends|final|implements|import|instanceof|interface|native|package|private|protected|public|static|strictfp|super|synchronized|throws|volatile|sealed)\\b"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant: "\\b(true|false|null)\\b"
    - comment:
        start: "//"
        end: "$"
        rules: []
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules: []
    - comment:
        start: "/\\*\\*"
        end: "\\*/"
        rules: []`)
	}}
}
