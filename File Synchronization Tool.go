package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func syncDirectories(source, destination string) {
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(source, path)
		destPath := filepath.Join(destination, relPath)

		if info.IsDir() {
			os.MkdirAll(destPath, info.Mode())
		} else {
			copyFile(path, destPath)
		}

		return nil
	})
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = sourceFile.WriteTo(destFile)
	return err
}

func main() {
	sourceDir := "/path/to/source/directory"
	destinationDir := "/path/to/destination/directory"

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		syncDirectories(sourceDir, destinationDir)
	}()

	wg.Wait()
	fmt.Println("Synchronization completed.")
}
