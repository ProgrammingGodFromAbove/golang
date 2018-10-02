package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)
	m["k1"] = 8
	m["k3"] = 21

	fmt.Println(m)

	a, b := m["k3"]

	fmt.Println(a, b)

}
