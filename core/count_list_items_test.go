package core

import (
	"os"
	"testing"
)

func TestCountListItems(t *testing.T) {
	input, err := os.ReadFile("../testdata/test.cue")
	if err != nil {
		t.Fatalf("Error reading test file: %v", err)
	}

	filelistCount, err := CountListItems(string(input), "filelist")
	if err != nil {
		t.Fatalf("Error counting filelist items: %v", err)
	}
	if filelistCount != 2 {
		t.Errorf("Expected 2 items in filelist, got %d", filelistCount)
	}

	filelist2Count, err := CountListItems(string(input), "filelist2")
	if err != nil {
		t.Fatalf("Error counting filelist2 items: %v", err)
	}
	if filelist2Count != 2 {
		t.Errorf("Expected 2 items in filelist2 (one commented out), got %d", filelist2Count)
	}
}
