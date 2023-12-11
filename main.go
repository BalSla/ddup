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

func main() {

	flag.Parse()

	if *sourceDir == "" || *targetDir == "" {
		fmt.Println("Source and target directories are required")
		os.Exit(1)
	}

	sourceDir := *sourceDir
	targetDir := *targetDir

	// Build a map of filenames from the source directory
	sourceFiles := make(map[string]bool)
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
			// fmt.Printf("Found file: %s\n", relPath)
			sourceFiles[relPath] = true
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

		if !info.IsDir() {
			relPath, err := filepath.Rel(targetDir, path)
			if err != nil {
				return err
			}
			if _, exists := sourceFiles[relPath]; exists {
				// fmt.Printf("Found file: %s\n", relPath)
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
