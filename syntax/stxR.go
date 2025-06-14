//go:build stxR || stxAll

package syntax

func init() {
	syntaxMap["r"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: r
rules:
    - statement: "\\b(break|else|for|function|if|in|next|repeat|return|while)\\b"
    - constant: "\\b(TRUE|FALSE|NULL|Inf|NaN|NA|NA_integer_|NA_real_|NA_complex_|NA_character_)\\b"
    - comment:
        start: "#"
        end: "$"
        rules: []`)
	}}
}
