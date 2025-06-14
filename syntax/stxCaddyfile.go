//go:build stxCaddyfile || stxAll

package syntax

func init() {
	syntaxMap["caddyfile"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: caddyfile
rules:
    - identifier: "^\\s*\\S+(\\s|$)"
    - type: "^([\\w.:/-]+,? ?)+[,{]$"
    - constant.specialChar: "\\s{$"
    - constant.specialChar: "^\\s*}$"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - preproc: "\\{(\\w+|\\$\\w+|%\\w+%)\\}"
    - comment:
        start: "#"
        end: "$"
        rules: []`)
	}}
}
