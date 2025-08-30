package intermediate

import "os"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func manageDIr() {
	// checkError(os.Mkdir("new-folder", 0755))
	// defer checkError(os.RemoveAll("new-directory"))

	// checkError(os.WriteFile("new-folder/file", []byte{}, 0755))

	checkError(os.MkdirAll("subdir/parent/child", 0755))
}