package utils

import "testing"

func TestGenerateBaseFilenameWithoutASIN(t *testing.T) {
	if got := GenerateBaseFilename("Title", "Subtitle", ""); got != "Title: Subtitle" {
		t.Fatalf("got %q", got)
	}
}
