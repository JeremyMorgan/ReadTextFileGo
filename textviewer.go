package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	file, ferr := os.Open("customerlist.csv")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		fmt.Printf("Name: %s %s Email: %s\n", items[1],items[2],items[3])
		fmt.Println("------")
	}
}