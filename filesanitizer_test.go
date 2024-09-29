package filesanitizer

import (
	"testing"
)

func TestSanitizeFilesInFolder(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		err := SanitizeFilesInFolder("./test")

		if err != nil {
			t.Fatal(err.Error())
		}
	})
}

func TestConvertFileName(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		expected := "foo_bar.txt"
		result := convertFileName("foo & ba)r..txt")

		if expected != result {
			t.Fatalf("File name is wrong. Expected: %s but was: %s", expected, result)
		}
	})
}
