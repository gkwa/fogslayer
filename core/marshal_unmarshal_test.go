package core

import (
	"os"
	"strings"
	"testing"
)

func TestMarshalUnmarshalCUE(t *testing.T) {
	input, err := os.ReadFile("../testdata/test.cue")
	if err != nil {
		t.Fatalf("Error reading test file: %v", err)
	}

	result, err := MarshalUnmarshalCUE(string(input))
	if err != nil {
		t.Fatalf("Error in MarshalUnmarshalCUE: %v", err)
	}

	if !strings.Contains(result, "// This is a test comment") {
		t.Error("Comment was not preserved")
	}

	if !strings.Contains(result, `b: "updated"`) {
		t.Error("Value was not updated")
	}

	if !strings.Contains(result, `a: "test"`) {
		t.Error("Unmodified value was not preserved")
	}

	if !strings.Contains(result, "filelist: [") {
		t.Error("List was not preserved")
	}

	if !strings.Contains(result, `"/path1"`) || !strings.Contains(result, `"/path2"`) {
		t.Error("List items were not preserved")
	}

	t.Logf("Result:\n%s", result)
}
