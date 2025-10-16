package main

import (
	"dz2/internal/task_1"
	"strings"
	"fmt"
)

func main()	{
	input := `q--3wer4301-4+r_e**
	end`
	results := task1.AnimalFeeding(strings.NewReader(input))
	fmt.Println(results)
}
