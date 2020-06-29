package internal

import (
	"fmt"
	"log"
	"os"
)

func PrintDir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
