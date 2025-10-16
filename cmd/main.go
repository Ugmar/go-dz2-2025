package main

import (
	"dz2/internal/task_1"
	"strings"
	"fmt"
)

func main()	{
	input := `q--3wer4301-4+r_e**
	dog d2
	cat c1
	cat c2
	bird b1
	end`
	results := task1.AnimalFeeding(strings.NewReader(input))
	fmt.Println(results)
}
