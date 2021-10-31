package main

import (
	"fmt"

	"github.com/bugangongwei/go_training_camp/week02"
)

func main() {
	// record not found
	fmt.Println(week02.GetAuthor(5))

	// record exist
	fmt.Println(week02.GetAuthor(1))
}
