//go:build stxD || stxAll

package lexers

func init() {
	syntaxMap["d"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: d
rules:
    - statement: "(\\*|/|%|\\+|-|>>|<<|>>>|&|\\^(\\^)?|\\||~)?="
    - statement: "\\.\\.(\\.)?|!|\\*|&|~|\\(|\\)|\\[|\\]|\\\\|/|\\+|-|%|<|>|\\?|:|;"
    - error: "(0[0-7_]*)(L[uU]?|[uU]L?)?"
    - constant.number: "([0-9]|[1-9][0-9_]*)(L[uU]?|[uU]L?)?"
    - constant: "(0[bB][01_]*)(L[uU]?|[uU]L?)?"
    - constant.number: "[0-9][0-9_]*\\.([0-9][0-9_]*)([eE][+-]?([0-9][0-9_]*))?[fFL]?i?"
    - constant.number: "[0-9][0-9_]*([eE][+-]?([0-9][0-9_]*))[fFL]?i?"
    - constant.number: "[^.]\\.([0-9][0-9_]*)([eE][+-]?([0-9][0-9_]*))?[fFL]?i?"
    - constant.number: "[0-9][0-9_]*([fFL]?i|[fF])"
    - constant.number: "(0[xX]([0-9a-fA-F][0-9a-fA-F_]*|[0-9a-fA-F_]*[0-9a-fA-F]))(L[uU]?|[uU]L?)?"
    - constant.number: "0[xX]([0-9a-fA-F][0-9a-fA-F_]*|[0-9a-fA-F_]*[0-9a-fA-F])(\\.[0-9a-fA-F][0-9a-fA-F_]*|[0-9a-fA-F_]*[0-9a-fA-F])?[pP][+-]?([0-9][0-9_]*)[fFL]?i?"
    - constant.number: "0[xX]\\.([0-9a-fA-F][0-9a-fA-F_]*|[0-9a-fA-F_]*[0-9a-fA-F])[pP][+-]?([0-9][0-9_]*)[fFL]?i?"
    - constant.string:
        start: "'"
        end: "'"
        rules:
            - constant.specialChar: "\\\\."
    - statement: "\\b(abstract|alias|align|asm|assert|auto|body|break|case|cast|catch|class|const|continue|debug|default|delegate|do|else|enum|export|extern)\\b"
    - statement: "\\b(false|final|finally|for|foreach|foreach_reverse|function|goto|if|immutable|import|in|inout|interface|invariant|is|lazy)\\b"
    - statement: "\\b(macro|mixin|module|new|nothrow|null|out|override|package|pragma|private|protected|public|pure|ref|return)\\b"
    - statement: "\\b(scope|shared|static|struct|super|switch|synchronized|template|this|throw|true|try|typeid|typeof|union|unittest|version|while|with)\\b"
    - statement: "\\b(__FILE__|__MODULE__|__LINE__|__FUNCTION__|__PRETTY_FUNCTION__|__gshared|__traits|__vector|__parameters)\\b"
    - error: "\\b(delete|deprecated|typedef|volatile)\\b"
    - type: "\\b(bool|byte|cdouble|cent|cfloat|char|creal|dchar|double|float|idouble|ifloat|int|ireal|long|real|short|ubyte|ucent|uint|ulong|ushort|void|wchar)\\b"
    - type: "\\b(string|wstring|dstring|size_t|ptrdiff_t)\\b"
    - constant: "\\b(__DATE__|__EOF__|__TIME__|__TIMESTAMP__|__VENDOR__|__VERSION__)\\b"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "r\""
        end: "\""
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "` + "`" + `"
        end: "` + "`" + `"
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "x\""
        end: "\""
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "q\"\\("
        end: "\\)\""
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "q\"\\{"
        end: "q\"\\}"
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "q\"\\["
        end: "q\"\\]"
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "q\"<"
        end: "q\">"
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "q\"[^({[<\"][^\"]*$"
        end: "^[^\"]+\""
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "q\"([^({[<\"])"
        end: "\""
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: "//"
        end: "$"
        rules: []
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules: []
    - comment:
        start: "/\\+"
        end: "\\+/"
        rules: []`)
	}}
}
