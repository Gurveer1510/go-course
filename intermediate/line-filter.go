package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func line_filter() {

	file, _ := os.Open("example.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// keyword to filter line
	keyword := "important"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updated_line := strings.ReplaceAll(line, keyword, "Necessary")
			fmt.Println("Filtered line",line)
			fmt.Println("Replaced line",updated_line)
		}
	}

	err := scanner.Err()

	if err != nil {
		fmt.Println("Error scanning file:", err)
	}

}
