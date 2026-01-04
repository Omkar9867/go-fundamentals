package main

import (
	"fmt"
	"go-fundamentals/basics"
)

func main() {
	fmt.Println("=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('=') + "=" + string('='))
	fmt.Println()

	// Run basic examples
	fmt.Println("📚 BASIC CONCEPTS")
	fmt.Println("------------------")
	basics.RunVariablesExample()
	fmt.Println()
	basics.RunFunctionsExample()
	fmt.Println()
	basics.RunControlStructuresExample()
	fmt.Println()

}
