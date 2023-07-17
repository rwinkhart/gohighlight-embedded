package highlight

import (
        "bytes"
        "io/ioutil"
        "testing"

	"github.com/jessp01/gohighlight"
    )


func getDefs(t *testing.T, filename string, data []byte, highlight_lexer string) *highlight.Def {
	var def *highlight.Def
	syn_dir := "./syntax_files"

	var defs []*highlight.Def
	lerr := highlight.ParseSyntaxFiles(syn_dir, &defs)
	if lerr != nil {
		t.Errorf("Couldn't get defs from '%s', error: %v\n", syn_dir, lerr)
	}
	highlight.ResolveIncludes(defs)

	// Always try to auto detect the best lexer:was
	if def == nil {
		def = highlight.DetectFiletype(defs, filename, bytes.Split(data, []byte("\n"))[0])
	}

	// if a specific lexer was requested by setting the ENV var, try to load it
	if highlight_lexer != "" {
		syntaxFile, lerr := ioutil.ReadFile(syn_dir + "/" + highlight_lexer + ".yaml")
		if lerr == nil {
			// Parse it into a `*highlight.Def`
			def, _ = highlight.ParseDef(syntaxFile)
		}
	}

	if def == nil {
		t.Errorf("Found no defs for '%s'\n", filename)
	}

	//fmt.Println(def)
	return def
}
