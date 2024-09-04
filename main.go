package main

import (
	"fmt"

	"go.mod/config"
)

func main() {
	fmt.Println("Niggaaa")

	cfg, err := config.NewConfig()

	fmt.Println(cfg)
	fmt.Println(err)
}
