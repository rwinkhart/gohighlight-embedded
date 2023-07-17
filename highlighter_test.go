package highlight

import (
	"io/ioutil"
        "testing"
	"fmt"
	"path/filepath"
	"strings"
)

func TestInputs(t *testing.T){

    test_inputs_dir := "./test_inputs/*"
    files, err := filepath.Glob(test_inputs_dir)
    if err != nil {
	t.Errorf("Couldn't open '%s', error: %v\n", test_inputs_dir, err)
    }
    for _, filename := range(files){
	ext := strings.Split(filename, ".")
	data, _ := ioutil.ReadFile(filename)
	fmt.Println("Testing def detection for " + filename)
	def := getDefs(t, filename, data, "")
	exp := ext[len(ext) -1]
	if def.FileType != exp {
	    t.Errorf("\nInput   [%#v]\nExpected[%#v]\nGot     [%#v]\n",
	    //string(data), exp, def.FileType)
	    filename, exp, def.FileType)
	}
   }
}
