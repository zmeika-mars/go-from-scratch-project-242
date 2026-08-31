package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize returns the size of a file or directory in the requested format.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

// getSize calculates the size of a file or directory.
func getSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	// If it's a file, we check whether it's hidden or not
	if !info.IsDir() {
		return info.Size(), nil
	}

	//If a directory, get its contents
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	//Variable for storing the total sum of the dimensions
	var rawSize int64

	for _, file := range files {
		//Skip hidden files and directories unless --all is specified.
		if strings.HasPrefix(file.Name(), ".") && !all {
			continue
		}

		//If the current item is a directory
		if file.IsDir() {
			// If the ‘-r’ flag is specified, check the files inside the directories
			if recursive {
				size, err := getSize(filepath.Join(path, file.Name()), recursive, all)
				if err != nil {
					//Skip unreadable directories and continue.
					continue
				}

				//Add the file size to the total
				rawSize += size
			}
			continue
		}
		//If the current item is a file, we retrieve information about it
		info, err := file.Info()
		if err != nil {
			//Skip unreadable directories and continue.
			continue
		}
		//Add the file size to the total size.
		rawSize += info.Size()
	}
	return rawSize, nil
}

// formatSize formats a size in bytes as a human-readable string when requested.
func formatSize(size int64, human bool) string {
	//If the --human flag is not passed, return the size in bytes
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	//List of Units of Measurement
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

	value := float64(size)
	//Index of the Current Unit of Measurement
	unit := 0

	for value >= 1024 && unit < len(units)-1 {
		//Divide by 1024 and move on to the next unit of measurement as long as the value is still greater than or equal to 1024
		value = value / 1024
		unit++
	}
	//If the size is less than 1024 bytes, display the size in bytes
	if unit == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unit])
}
