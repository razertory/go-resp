package server

import "fmt"

func debug(str string, args ...interface{}) {
	fmt.Println(str, args)
}
