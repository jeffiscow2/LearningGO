package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your NAME:")
	var name string
	fmt.Scan(&name)
	fmt.Fprintln(os.Stdout, []any{"Hello ", name}...)
}
