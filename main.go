package main

import (
	"fmt"
	"log"
	"module/app"
	"os"
)

func main() {
	fmt.Println("Starting point")

	application := app.Generate()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
