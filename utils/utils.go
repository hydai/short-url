package utils

import "log"

func ErrorOrNil(e error) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("Handle panic ", r)
		}
	}()

	if e != nil {
		panic(e)
	}
}
