package metadata

import "testing"

func TestIsEndMarker(t *testing.T) {
	for _, title := range []string{"End", "end", "- End", "-- End"} {
		if !isEndMarker(title) {
			t.Errorf("isEndMarker(%q) = false, want true", title)
		}
	}
	for _, title := range []string{"Ending", "The End", "Chapter 1"} {
		if isEndMarker(title) {
			t.Errorf("isEndMarker(%q) = true, want false", title)
		}
	}
}
