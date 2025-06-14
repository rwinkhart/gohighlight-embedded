//go:build stxRuby || stxAll

package lexers

func init() {
	syntaxMap["ruby"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: ruby
rules:
    - statement: "\\b(BEGIN|END|alias|and|begin|break|case|class|def|defined\\?|do|else|elsif|end|ensure|false|for|if|in|module|next|nil|not|or|redo|rescue|retry|return|self|super|then|true|undef|unless|until|when|while|yield)\\b"
    - constant: "(\\$|@|@@)?\\b[A-Z]+[0-9A-Z_a-z]*"
    - constant.number: "\\b[0-9]+\\b"
    - constant: "(i?)([ 	]|^):[0-9A-Z_]+\\b"
    - constant: "\\b(__FILE__|__LINE__)\\b"
    - constant: "/([^/]|(\\\\/))*/[iomx]*|%r\\{([^}]|(\\\\}))*\\}[iomx]*"
    - constant.string: "` + "`" + `[^` + "`" + `]*` + "`" + `|%x\\{[^}]*\\}"
    - constant.string: "\"([^\"]|(\\\\\"))*\"|%[QW]?\\{[^}]*\\}|%[QW]?\\([^)]*\\)|%[QW]?<[^>]*>|%[QW]?\\[[^]]*\\]|%[QW]?\\$[^$]*\\$|%[QW]?\\^[^^]*\\^|%[QW]?![^!]*!"
    - special: "#\\{[^}]*\\}"
    - constant.string: "'([^']|(\\\\'))*'|%[qw]\\{[^}]*\\}|%[qw]\\([^)]*\\)|%[qw]<[^>]*>|%[qw]\\[[^]]*\\]|%[qw]\\$[^$]*\\$|%[qw]\\^[^^]*\\^|%[qw]![^!]*!"
    - comment: "#+[^{].*$|#$"
    - comment:
        start: "=begin"
        end: "=end"
        rules: []
    - constant.macro:
        start: "<<-?'?EOT'?"
        end: "^EOT"
        rules: []
    - todo: "(XXX|TODO|FIXME|\\?\\?\\?)"`)
	}}
}
