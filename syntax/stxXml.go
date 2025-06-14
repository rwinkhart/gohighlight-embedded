//go:build stxXml || stxAll

package syntax

func init() {
	syntaxMap["xml"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: xml
rules:
    - type: "(<(/)?[a-zA-Z:]+>)"
    - type: "(<(/)?[a-zA-Z:]+[[:space:]]+)"
    - constant.specialChar: '[a-zA-Z0-9]+="(.*)+"'
    - statement: '([a-zA-Z0-9]+)='
    - type: "(([[:space:]])?/?>([[:space:]]+)?$)"
    - comment:
        start: "<!DOCTYPE"
        end: "[/]?>"
        rules: []
    - comment:
        start: "<!--"
        end: "-->"
        rules: []
    - special: "&[^;]*;"`)
	}}
}
