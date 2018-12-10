package app

import "fmt"

type Tracker struct {
}

func (tracker Tracker) Run() {
	fmt.Println("running")
}
