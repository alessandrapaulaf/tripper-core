package main

import (
	"fmt"

	"github.com/alessandrapaulaf/tripper-core/configs"
)

func main() {
	configs.StartDb()
	fmt.Println("Hello World")
}
