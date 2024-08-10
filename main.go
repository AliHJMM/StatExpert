package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}

	filename := os.Args[1]
}