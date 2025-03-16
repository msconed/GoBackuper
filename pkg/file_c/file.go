package file_c

import (
	"fmt"
	"os"
	"path/filepath"
)

func CountFilesInDirectory(dir string) (int, error) {
	var fileCount int
	stack := []string{dir}
	for len(stack) > 0 {
		currentDir := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		entries, err := os.ReadDir(currentDir)
		if err != nil {
			return 0, err
		}
		for _, entry := range entries {
			path := filepath.Join(currentDir, entry.Name())
			if entry.IsDir() {
				stack = append(stack, path)
			} else {
				fileCount++
			}
		}
	}
	return fileCount, nil
}

func RemoveFileIfExist(file string) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		defer func() {
			if err := os.Remove(file); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("File deleted successfully.")
			}
		}()
	} else {
		fmt.Println("File does not exist.")
	}
}
