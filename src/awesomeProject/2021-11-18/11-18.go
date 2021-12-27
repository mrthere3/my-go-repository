package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Hello世界！123Go."
	reg := regexp.MustCompile(`[^a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}
