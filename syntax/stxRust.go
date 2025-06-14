//go:build stxRust || stxAll

package syntax

func init() {
	syntaxMap["rust"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: rust
rules:
    - identifier: "fn [a-z0-9_]+"
    - statement: "\\b(abstract|alignof|as|become|box|break|const|continue|crate|do|else|enum|extern|false|final|fn|for|if|impl|in|let|loop|macro|match|mod|move|mut|offsetof|override|priv|pub|pure|ref|return|sizeof|static|self|struct|super|true|trait|type|typeof|unsafe|unsized|use|virtual|where|while|yield)\\b"
    - special: "[a-z_]+!"
    - constant: "[A-Z][A-Z_]+"
    - constant.number: "\\b[0-9]+\\b"
    - type: "[A-Z][a-z]+"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "r#+\""
        end: "\"#+"
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
            - todo: "(TODO|XXX|FIXME):?"
    - special:
        start: "#!\\["
        end: "\\]"
        rules: []`)
	}}
}
