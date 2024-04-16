package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	name       string
	age        string
	occupation string
}

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var user []User
	scanner := bufio.NewScanner(f)
	var person User
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ":")
		// fmt.Println(words[0])
		// fmt.Println(words[1])
		switch words[0] {
		case "Name":
			person.name = words[1]
		case "Age":
			person.age = words[1]
		case "Occupation":
			person.occupation = words[1]
			user = append(user, person)
		}
	}
	for _, i := range user {
		fmt.Println(i)
	}
}
