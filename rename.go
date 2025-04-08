package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	re := regexp.MustCompile(`^lesson(\d+)\.vtt$`)

	for _, file := range files {
		oldName := file.Name()

		match := re.FindStringSubmatch(oldName)
		if match != nil {
			number := match[1]
			newName := number + ".vtt"

			err := os.Rename(oldName, newName)
			if err != nil {
				fmt.Println("Error renaming", oldName, "to", newName, ":", err)
			} else {
				fmt.Println("Renamed", oldName, "->", newName)
			}
		}
	}
}
