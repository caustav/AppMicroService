package comp1

import "fmt"

type Trigger struct {
}

func (trigger Trigger) Submit() {
	fmt.Println("comp1 trigger1 printed")
}
