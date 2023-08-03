package highlight

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestInputs(t *testing.T) {

	testInputsDir := "./test_inputs/*"
	files, err := filepath.Glob(testInputsDir)
	if err != nil {
		t.Errorf("Couldn't open '%s', error: %v\n", testInputsDir, err)
	}
	for _, filename := range files {
		ext := strings.Split(filename, ".")
		data, _ := ioutil.ReadFile(filename)
		fmt.Println("Testing def detection for " + filename)
		def := getDefs(t, filename, data, "")
		exp := ext[len(ext)-1]
		if def.FileType != exp {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nGot     [%#v]\n",
				// string(data), exp, def.FileType)
				filename, exp, def.FileType)
		}
	}
}
