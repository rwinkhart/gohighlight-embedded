//go:build stxDockerfile || stxAll

package lexers

func init() {
	syntaxMap["dockerfile"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: dockerfile
rules:
    - keyword: "(?i)^(FROM|MAINTAINER|RUN|CMD|LABEL|EXPOSE|ENV|ADD|COPY|ENTRYPOINT|VOLUME|USER|WORKDIR|ONBUILD|ARG|HEALTHCHECK|STOPSIGNAL|SHELL)[[:space:]]"
    - statement: "(\\(|\\)|\\[|\\])"
    - special: "&&"
    - comment:
        start: "#"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
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
            - constant.specialChar: "\\\\."`)
	}}
}
