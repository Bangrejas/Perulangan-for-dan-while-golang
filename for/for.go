package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 50; i*=2 {
		fmt.Println("Nomor",i)
	}
}