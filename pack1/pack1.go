package pack1

import (
	"fmt"
	"github.com/google/uuid"
)

func P() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)
}