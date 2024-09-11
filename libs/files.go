package libs

import (
	"os"
	"strings"
)

func ListDirectory(dir string) ([]string, int, error) {
    // Open the directory
    entries, err := os.ReadDir(dir)
    if err != nil {
        return []string{}, 0, err
    }

	// Filter out hidden files (Unix-like systems)
	countDirs := 0
    filteredDirs := []string{}
	filteredDirs = append(filteredDirs, "..")
    for _, entry := range entries {
		// Exclude hidden files
        if !strings.HasPrefix(entry.Name(), ".") {
			// Append dir and file different
			if (entry.IsDir()) {
            	filteredDirs = append(filteredDirs, entry.Name() + "/")
			} else {
            	filteredDirs = append(filteredDirs, entry.Name())
			}
			countDirs += 1
        }
    }

    return filteredDirs, countDirs, nil
}
