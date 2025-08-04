package basics

import (
	"fmt"
	foo "net/http"
)

func ImportTut () {
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
}