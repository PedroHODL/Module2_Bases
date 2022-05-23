package main

import (
	"fmt"
	"os"
)

func main() {
	_ = os.Setenv("Name", "PedroHODL")
	value, ok := os.LookupEnv("Namae")
	fmt.Println(value, ok)
}
