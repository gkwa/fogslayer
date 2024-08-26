package core

import (
	"os"
	"testing"
)

func TestIsItemInList(t *testing.T) {
	input, err := os.ReadFile("../testdata/test.cue")
	if err != nil {
		t.Fatalf("Error reading test file: %v", err)
	}

	expectedItems := []string{"/path1", "/path5"}
	unexpectedItem := "/path3"

	for _, item := range expectedItems {
		inList, err := IsItemInList(string(input), "filelist2", item)
		if err != nil {
			t.Fatalf("Error checking if %s is in filelist2: %v", item, err)
		}
		if !inList {
			t.Errorf("Expected item %s to be in filelist2, but it was not found", item)
		}
	}

	inList, err := IsItemInList(string(input), "filelist2", unexpectedItem)
	if err != nil {
		t.Fatalf("Error checking if %s is in filelist2: %v", unexpectedItem, err)
	}
	if inList {
		t.Errorf("Unexpected item %s found in filelist2", unexpectedItem)
	}
}
