package main

import (
	"fmt"
	"github.com/q946666800/gomodule/pack1"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)
	pack1.P()
}
