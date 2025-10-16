package task1

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type AnimalAction interface{
	Eat(c rune)
	WhatDidYouEat() string
	GetName() string
}

type Animal struct {
	food string
	name string
}

type Cat struct{
	Animal
}

type Dog struct{
	Animal
}

type Bird struct{
	Animal
}

func (cat *Cat) Eat(c rune) {
	if unicode.IsDigit(c){
		cat.food += string(c)
	}
}

func (dog *Dog) Eat(c rune) {
	if unicode.IsLetter(c){
		dog.food += string(c)
	}
}

func (bird *Bird) Eat(c rune) {
	if !unicode.IsLetter(c) && !unicode.IsDigit(c){
		bird.food += string(c)
	}
}

func (animal Animal) WhatDidYouEat()string {
	return animal.food
}

func (animal Animal) GetName() string{
	return animal.name
}

func AnimalFeeding(stream io.Reader) ([]string){
	scanner := bufio.NewScanner(stream)

	scanner.Scan()
	feeder := scanner.Text()

	var query []AnimalAction

	scanner.Scan()
	line := strings.TrimLeft(scanner.Text(), "\t ")

	for line != "end"{
		elems := strings.Split(line, " ")
		kind := elems[0]
		name := elems[1]

		switch strings.ToLower(kind){
		case "cat":
			query = append(query, &Cat{Animal: Animal{name: name}})
		case "dog":
			query = append(query, &Dog{Animal: Animal{name: name}})
		case "bird":
			query = append(query, &Bird{Animal: Animal{name: name}})
		}

		scanner.Scan()
		line = strings.TrimLeft(scanner.Text(), "\t ")
	}

	for i, food := range(feeder){
		animal := query[i % len(query)]
		animal.Eat(food)
	}

	results := make([]string, len(query))
	for i, animal := range(query){
		results[i] = fmt.Sprintf("%s %s",  animal.GetName(), animal.WhatDidYouEat())
	}

	return results
}
