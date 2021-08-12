package pack1

import (
	"fmt"
	"github.com/google/uuid"
)

// P Ptest
func P() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)
}
