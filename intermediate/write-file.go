package intermediate

import (
	"fmt"
	"os"
)

func write_file() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}

	defer file.Close()

	// write data to file

	data := []byte("Hello world ! \n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
	}

	newFile, err := os.Create("new_file.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer newFile.Close()

	_, err = newFile.WriteString("Hello new file\n Second line")
	if err != nil {
		fmt.Println("Error writing string to newfile:", err)
		return
	}

	fmt.Println("Data written to new-file")

}
