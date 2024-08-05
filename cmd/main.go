package main

import (
	"fmt"
	"github.com/itpourya/golang-persian-tools/commas"
)

func main() {
	r := commas.AddCommas("300000")
	fmt.Println(r)
}
