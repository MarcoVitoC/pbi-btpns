package main

import (
	"github.com/MarcoVitoC/pbi-btpns/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}