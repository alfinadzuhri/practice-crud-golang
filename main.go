package main

import (
	"fmt"

	"github.com/alfinadzuhri/practice-crud-golang/config"
)

func main() {
	fmt.Println("check")

	cfg, err := config.NewConfig()

	fmt.Println(cfg)
	fmt.Println(err)
}
