package printLog

import (
	"fmt"
	"strings"
)

func Print(header string, body string) {
	fmt.Print(strings.Repeat("-", 50))

	fmt.Print("LOG INFO")
	fmt.Println(strings.Repeat("-", 42))

	fmt.Println("File: ./logs/...")

	fmt.Print(strings.Repeat("-", 50))
	fmt.Print("HEADER")
	fmt.Println(strings.Repeat("-", 44))

	fmt.Println(header)

	fmt.Print(strings.Repeat("-", 50))
	fmt.Print("BODY")
	fmt.Println(strings.Repeat("-", 46))

	fmt.Println(body)

	fmt.Println(strings.Repeat("-", 100))
}
