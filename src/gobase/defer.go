package gobase

import (
	"log"
)

func Log() {
	log.Println("start")
	defer log.Println("defer1")
	log.Println("start2")
	defer log.Println("defer2")
	log.Println("start3")
	defer log.Println("defer3")
}
