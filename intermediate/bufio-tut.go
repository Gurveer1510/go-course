package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func bufio_tut() {
	// reader := bufio.NewReader(strings.NewReader("Hello bufio package\nhooray"))

	// data := make([]byte, 20)
	// n, err := reader.Read(data)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Bytes read:", n)
	// fmt.Printf("Read string: %s", data[:n])

	// line, err := reader.ReadString('\n')

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Line:", line)

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data := []byte("hello bufio package!\n")
	n, err := writer.Write(data)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Wrote %d bytes\n", n)
	writer.Flush()
}
