package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func file_path_tut() {

	relative_path := "./data/file.txt"
	absolute_path := "C:/Users/pwar-qa-4/go/go-course/example.txt"

	joinedPath := filepath.Join("downloads", "file.zip")
	fmt.Println("Joined Path", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("normalized path", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Dir", dir)
	fmt.Println("file", file)

	base:= filepath.Base("/home/user/docs/file.txt")
	fmt.Println("Base", base)

	fmt.Println("Is relativePath variable absolute", filepath.IsAbs(relative_path))
	fmt.Println("Is absolutePath variable absolute", filepath.IsAbs(absolute_path))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/title.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/title.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rel)

	abs, err := filepath.Abs(relative_path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(abs)
}

