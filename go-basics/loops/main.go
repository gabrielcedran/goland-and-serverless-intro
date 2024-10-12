package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := [2]string{}

	animals[0] = "dog"
	animals[1] = "cat"

	fmt.Println(animals)

	animals2 := []string{"dog", "cat"}
	animals2 = append(animals2, "bird")

	animals2 = slices.Delete(animals2, 1, 2)
	// animals2 = append(animals2[0:1], animals2[2:3]...)

	fmt.Println(animals2)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for index, animal := range animals2 {
		fmt.Printf("My index %d, my animal %s\n", index, animal)
	}

	for index := range 10 {
		fmt.Println(index)
	}
}
