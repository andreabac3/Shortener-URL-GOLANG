package main

import (
	L "./lib"
	"fmt"
)

func main() {

	fmt.Println("Start shorturl in pure golang")
	if L.Start_Server() < 0 {
		defer L.EndDB(L.Client)
	}

	fmt.Println("End shorturl in pure golang")
}
