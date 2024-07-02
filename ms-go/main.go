package main

import (
	"ms-go/router"
)

func main() {
	go func() {
		router.Run()
	}()
	
	select {}
}
