//go:build stxJs || stxAll

package lexers

func init() {
	syntaxMap["js"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: javascript
rules:
    - type: "(<(/)?[a-zA-Z:]+>)"
    - type: "(<(/)?[a-zA-Z:]+[[:space:]]+)"
    - constant.number: "\\b[-+]?([1-9][0-9]*|0[0-7]*|0x[0-9a-fA-F]+)([uU][lL]?|[lL][uU]?)?\\b"
    - constant.number: "\\b[-+]?([0-9]+\\.[0-9]*|[0-9]*\\.[0-9]+)([EePp][+-]?[0-9]+)?[fFlL]?"
    - constant.number: "\\b[-+]?([0-9]+[EePp][+-]?[0-9]+)[fFlL]?"
    - identifier: "[A-Za-z_][A-Za-z0-9_]*[[:space:]]*[(]"
    - special: "[\\(|\\)|/|+|\\-|\\*|\\[|.|,|;]"
    - statement: "\\b(break|case|catch|continue|default|delete|do|else|finally)\\b"
    - statement: "\\b(for|function|get|if|in|instanceof|new|return|set|switch)\\b"
    - statement: "\\b(switch|this|throw|try|typeof|var|void|while|with|const|let)\\b"
    - constant: "\\b(null|undefined|NaN)\\b"
    - constant: "\\b(true|false)\\b"
    - type: "\\b(Array|Boolean|Date|Enumerator|Error|Function|Math)\\b"
    - type: "\\b(Number|Object|RegExp|String|require|import)\\b"
    - statement: "[-+/*=<>!~%?:&|]"
    - constant: "/[^*]([^/]|(\\\\/))*[^\\\\]/[gim]*"
    - constant: "\\\\[0-7][0-7]?[0-7]?|\\\\x[0-9a-fA-F]+|\\\\[bfnrt'\"\\?\\\\]"
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
    - comment:
        start: "//"
        end: "$"
        rules: []
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules: []`)
	}}
}
