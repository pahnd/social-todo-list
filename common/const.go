package common

import "fmt"

const (
	CurrentUser = "current_user"
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}
