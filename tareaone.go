package main

import (
	"fmt"
)

func main() {

	coco := "La Cancion Del Amor"

	for luis := 0; luis < len(coco); luis++ {

		fmt.Println(luis, string(coco[luis]))
	}

}
