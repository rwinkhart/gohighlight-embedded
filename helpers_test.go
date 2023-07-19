package highlight

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func getDefs(t *testing.T, filename string, data []byte, highlight_lexer string) *Def {
	var def *Def
    syn_dir := "./syntax_files"

	var defs []*Def
	lerr, warnings := ParseSyntaxFiles(syn_dir, &defs)
	if lerr != nil {
		t.Errorf("Couldn't get defs from '%s', error: %v\n", syn_dir, lerr)
	}
	if len(warnings) > 0 {
		t.Errorf("Parsing ended with warnings: '%s'\n", strings.Join(warnings, ";"))
	}

	ResolveIncludes(defs)

	// Always try to auto detect the best lexer:was
	if def == nil {
		def = DetectFiletype(defs, filename, bytes.Split(data, []byte("\n"))[0])
	}

	// if a specific lexer was requested by setting the ENV var, try to load it
	if highlight_lexer != "" {
		syntaxFile, lerr := ioutil.ReadFile(syn_dir + "/" + highlight_lexer + ".yaml")
		if lerr == nil {
			// Parse it into a `*highlight.Def`
			def, _ = ParseDef(syntaxFile)
		}
	}

	if def == nil {
		t.Errorf("Found no defs for '%s'\n", filename)
	}

	//fmt.Println(def)
	return def
}
