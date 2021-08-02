package yamlplus

import (
	"strings"
	"testing"
)

var getTests = []struct {
	path []string
	line int
}{
	{[]string{"info"}, 1},
	{[]string{"info", "title"}, 2},
	{[]string{"info", "contact", "name"}, 6},
	{[]string{"tags"}, 12},
	{[]string{"tags", "1"}, 14},
	{[]string{"components"}, 30},
	{[]string{"servers"}, 21676},
	{[]string{"servers", "1", "description"}, 21680},
}

// TestGetNodeJsonSchemaPath validates `GetNodeJsonSchemaPath()`.
func TestGetNodeJsonSchemaPath(t *testing.T) {
	testfile := "testdata/openapi3_spec.yaml"
	topNode, err := ReadFile(testfile)
	if err != nil {
		t.Errorf("ParseFileNode(\"%s\") Error: [%s]", testfile, err.Error())
	}
	// topNodeThin := NodeToThin(topNode)
	// fmtutil.PrintJSON(topNodeThin)
	for _, tt := range getTests {
		tryNode, err := GetNodeJsonSchemaPath(topNode, tt.path...)
		if err != nil {
			t.Errorf("GetNodeJsonSchemaPath(\"%s\") Error: [%s]",
				strings.Join(tt.path, ","), err.Error())
		}
		if tryNode.Line != tt.line {
			t.Errorf("GetNodeJsonSchemaPath(\"%s\") Want Line: [%d] Got Line [%d]",
				strings.Join(tt.path, ","), tt.line, tryNode.Line)
		}
	}
}
