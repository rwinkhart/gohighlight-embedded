package highlight

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func getDefs(t *testing.T, filename string, data []byte, highlightLexer string) *Def {
	var def *Def
	synDir := "./syntax_files"

	var defs []*Def
	warnings, lerr := ParseSyntaxFiles(synDir, &defs)
	if lerr != nil {
		t.Errorf("Couldn't get defs from '%s', error: %v\n", synDir, lerr)
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
	if highlightLexer != "" {
		syntaxFile, lerr := ioutil.ReadFile(synDir + "/" + highlightLexer + ".yaml")
		if lerr == nil {
			// Parse it into a `*highlight.Def`
			def, _ = ParseDef(syntaxFile)
		}
	}

	if def == nil {
		t.Errorf("Found no defs for '%s'\n", filename)
	}

	return def
}
