//go:build stxGdscript || stxAll

package lexers

func init() {
	syntaxMap["gdscript"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: gdscript
rules:
    - constant: "\\b(null|self|true|false)\\b"
    - identifier: "\\b(abs|acos|asin|atan|atan2|ceil|clamp|convert|cos|cosh|db2linear|decimals|deg2rad|ease|exp|float|floor|fmod|fposmod|hash|int|isinf|isnan|lerp|linear2db|load|log|max|min|nearest_po2|pow|preload|print|printerr|printraw|prints|printt|rad2deg|rand_range|rand_seed|randomize|randi|randf|range|round|seed|sin|slerp|sqrt|str|str2var|tan|typeof|var2str|weakref)\\b"
    - identifier: "\\b(AnimationPlayer|AnimationTreePlayer|Button|Control|HTTPClient|HTTPRequest|Input|InputEvent|MainLoop|Node|Node2D|SceneTree|Spatial|SteamPeer|PacketPeer|PacketPeerUDP|Timer|Tween)\\b"
    - type: "\\b(Vector2|Vector3)\\b"
    - identifier: "func [a-zA-Z_0-9]+"
    - statement: "\\b(and|as|assert|break|breakpoint|class|const|continue|elif|else|export|extends|for|func|if|in|map|not|onready|or|pass|return|signal|var|while|yield)\\b"
    - special: "@.*[(]"
    - statement: "[.:;,+*|=!\\%@]|<|>|/|-|&"
    - statement: "[(){}]|\\[|\\]"
    - constant: "\\b[0-9]+\\b"
    - constant.number: "\\b([0-9]+|0x[0-9a-fA-F]*)\\b|'.'"
    - comment:
        start: "\"\"\""
        end: "\"\"\""
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - comment:
        start: "'''"
        end: "'''"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\([0-7]{3}|x[A-Fa-f0-9]{2}|u[A-Fa-f0-9]{4}|U[A-Fa-f0-9]{8})"
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\([0-7]{3}|x[A-Fa-f0-9]{2}|u[A-Fa-f0-9]{4}|U[A-Fa-f0-9]{8})"
    - constant.string:
        start: "` + "`" + `"
        end: "` + "`" + `"
        rules: []
    - comment:
        start: "#"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
