package main

import (
	"dz2/internal/task_1"
	"strings"
	"fmt"
)

func main()	{
	input := `a

	dog d1
	cat c1
	bird b1
	end`
	results := task1.AnimalFeeding(strings.NewReader(input))
	fmt.Println(results)
}
