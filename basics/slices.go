package basics

import "fmt"

func slices() {

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 1, 2}

	slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}

	slice2 := a[2:4]

	slice = append(slice, 10, 11, 12)

	fmt.Println(slice2)

	fmt.Println(slice)

	sliceCopy := make([]int, len(slice2))
	copy(sliceCopy, slice2)

	fmt.Println("Original slice: ", slice2)
	fmt.Println("Copied slice: ", sliceCopy)

	var nilSlice []int
	fmt.Println(nilSlice)

	for _, v := range slice {
		fmt.Println(v)
	}

	twD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twD[i] = make([]int, innerLen)
		for j := range innerLen {
			twD[i][j] = i + j
		}
	}

	fmt.Println(twD)
}
