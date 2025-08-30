package intermediate

import (
	"fmt"
	"os"
)

// func checkError2(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func temp_files_dir() {
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)

	fmt.Println("Temporary file created:", tempFile.Name())

	defer tempFile.Close()
	
	defer os.Remove(tempFile.Name())
}
