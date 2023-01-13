package main

import (
	"fmt"

	"github.com/wduartebr/goexpert/apis/configs"
)

func main() {

	cfg, _ := configs.LoadConfig(".")

	fmt.Print(cfg.DBDriver)
}
