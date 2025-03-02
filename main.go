package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: Change dryrun to false
var dryrun = flag.Bool("dryrun", false, "Dry run")
var sourceDir = flag.String("source", "", "Source directory")
var targetDir = flag.String("target", "", "Target directory")

// Add helper function to streamline duplicate check
func isDuplicate(relPath string, sourceFiles map[string]string) bool {
	// Check for an exact relative path match
	if _, exists := sourceFiles[relPath]; exists {
		return true
	}
	// Check if any base file name matches
	targetBase := filepath.Base(relPath)
	for _, base := range sourceFiles {
		if base == targetBase {
			return true
		}
	}
	return false
}

func main() {

	flag.Parse()

	if *sourceDir == "" || *targetDir == "" {
		fmt.Println("Source and target directories are required")
		os.Exit(1)
	}

	sourceDir := *sourceDir
	targetDir := *targetDir

	// Build a map of filenames from the source directory
	// Modified: using map[string]string where value is the base file name
	sourceFiles := make(map[string]string)
	fmt.Printf("Scanning source directory: %s\n", sourceDir)
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}
			// Store base file name as value
			sourceFiles[relPath] = filepath.Base(relPath)
		}
		return nil
	})

	coounter := 0
	fmt.Printf("Found %d files\n", len(sourceFiles))
	fmt.Printf("Scanning target directory: %s\n", targetDir)
	// Delete duplicate files in the target directory
	filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// TODO: Add '/MediaFiles' automatically
		if !info.IsDir() {
			relPath, err := filepath.Rel(targetDir, path)
			if err != nil {
				return err
			}
			if isDuplicate(relPath, sourceFiles) {
				coounter++
				if *dryrun {
					fmt.Printf("Would delete file: %s\n", path)
				} else {
					err := os.Remove(path)
					if err != nil {
						fmt.Printf("Failed to delete file: %s\n", path)
					} else {
						fmt.Printf("Deleted file: %s\n", path)
					}
				}
			}
		}
		return nil
	})
	fmt.Printf("Deleted %d files\n", coounter)
}
