package filesanitizer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func SanitizeFilesInFolder(dirPath string) error {
	dir := dirPath
	list, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range list {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		newFileName := convertFileName(name)
		logrus.Debug("OLD: ", name, " - NEW: ", newFileName)

		// rename the file
		os.Rename(filepath.Join(dir, name), filepath.Join(dir, newFileName))
	}

	return nil
}

func convertFileName(oldFileName string) string {
	newFileName := oldFileName
	extension := filepath.Ext(newFileName)

	// trim
	newFileName = strings.ReplaceAll(newFileName, "(", "")
	newFileName = strings.ReplaceAll(newFileName, ")", "")
	newFileName = strings.ReplaceAll(newFileName, "[", "")
	newFileName = strings.ReplaceAll(newFileName, "]", "")
	newFileName = strings.ReplaceAll(newFileName, "'", "")
	newFileName = strings.ReplaceAll(newFileName, "!", "")
	newFileName = strings.ReplaceAll(newFileName, "?", "")

	// map
	newFileName = strings.ReplaceAll(newFileName, "  ", "_")
	newFileName = strings.ReplaceAll(newFileName, " ", "_")
	newFileName = strings.ReplaceAll(newFileName, "_"+extension, extension)
	newFileName = strings.ReplaceAll(newFileName, "."+extension, extension)
	newFileName = strings.ReplaceAll(newFileName, "&", "-")
	newFileName = strings.ReplaceAll(newFileName, ",", ".")
	newFileName = strings.ReplaceAll(newFileName, "--", "-")
	newFileName = strings.ReplaceAll(newFileName, "-_", "_")
	newFileName = strings.ReplaceAll(newFileName, "._", "_")
	newFileName = strings.ReplaceAll(newFileName, "__", "_")

	return newFileName
}
