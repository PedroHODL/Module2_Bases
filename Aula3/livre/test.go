package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	_ = os.Setenv("Name", "PedroHODL")
	value, ok := os.LookupEnv("Namae")
	fmt.Println(value, ok)

	r := strings.NewReader("TESTANDO...")
	io.Copy(os.Stdout, r)
}
