//go:build stxToml || stxAll

package lexers

func init() {
	syntaxMap["toml"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: toml
rules:
    - statement: "(.*)[[:space:]]="
    - special: "="
    - special: "(\\[|\\])"
    - constant.number: "\\b([0-9]+|0x[0-9a-fA-F]*)\\b|'.'"
    - constant.number: "\\\\([0-7]{3}|x[A-Fa-f0-9]{2}|u[A-Fa-f0-9]{4}|U[A-Fa-f0-9]{8})"
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
    - constant.string:
        start: "` + "`" + `"
        end: "` + "`" + `"
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: "#"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
